package upgrade

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func UntargzCmd(file string, binDir string) error {
	absFile, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	cmd := exec.Command("tar", "-xzf", absFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = binDir
	return cmd.Run()
}

func ExtractTarGzFile(file string, targetDir string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return ExtractTarGz(f, targetDir)
}

func ExtractTarGz(gz io.Reader, targetDir string) error {
	uncompressedStream, err := gzip.NewReader(gz)
	if err != nil {
		return fmt.Errorf("gzip: %w", err)
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("tar read: %w", err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filepath.Join(targetDir, header.Name), 0755); err != nil {
				return fmt.Errorf("tar mkdir %s: %w", header.Name, err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(targetDir, header.Name))
			if err != nil {
				return fmt.Errorf("tar create file %s: %w", header.Name, err)
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("tar write file %s: %w", header.Name, err)
			}
		default:
			return fmt.Errorf("tar unrecognized type %v: %s", header.Typeflag, header.Name)
		}
	}
	return nil
}
