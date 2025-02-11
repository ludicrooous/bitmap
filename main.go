package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Represents BMP header structure (first 14 bytes)
type BMPHeader struct {
	FileType [2]byte // "BM"
	FileSize uint32  // File size in bytes
	Reserved uint32  // Reserved (always 0)
	// Reserved2  uint16  // Reserved (always 0)
	OffsetData uint32 // Offset to image data
}

// Represents DIB header structure (next 40 bytes)
type DIBHeader struct {
	DibHeaderSize uint32 // DIB Header size
	Width         int32  // Width of image in pixels
	Height        int32  // Height of image in pixels
	Planes        uint16 // Number of color planes (must be 1)
	BitCount      uint16 // Bits per pixel (e.g., 24 for true color)
	Compression   uint32 // Compression (0 for uncompressed)
	ImageSize     uint32 // Image size in bytes (can be 0 for uncompressed)
	XPixelsPerM   int32  // Horizontal resolution (pixels per meter)
	YPixelsPerM   int32  // Vertical resolution (pixels per meter)
	ColorsUsed    uint32 // Number of colors used (0 means all)
	ColorsImp     uint32 // Important colors (0 means all)
}

// Represents a single pixel in the image (for 24-bit BMP files)
type Pixel struct {
	Blue  byte
	Green byte
	Red   byte
}

// Represents a command-line option that consists of name and its value
type Option struct {
	Name  string // The option name (e.g., "--mirror", "--filter", "--rotate", etc)
	Value string // The associated value (e.g., "horizontal", "90", "negative", etc)
}

// Parses command-line arguments while maintaining order
func parseArgs(args []string) (command string, filename string, outputFilename string, orderedOptions []Option, err error) {
	if len(args) < 2 {
		return "", "", "", nil, errors.New("invalid number of arguments")
	}

	command = args[0] // "header" or "apply"

	// Handle "header" command (only requires filename)
	if command == "header" {
		if len(args) != 2 {
			return "", "", "", nil, errors.New("usage: ./bitmap header <bmp_file>")
		}
		filename = args[1]
		return command, filename, "", nil, nil
	}

	// Handle "apply" command (requires at least one option, input file, and output file)
	if command == "apply" {
		if len(args) < 4 {
			return "", "", "", nil, errors.New("usage: ./bitmap apply [options] <source_file> <output_file>")
		}

		filename = args[len(args)-2]       // Second-to-last argument is the source file
		outputFilename = args[len(args)-1] // Last argument is the output file

		for i := 1; i < len(args)-2; i++ { // Ignore the last two arguments (file names)
			if strings.HasPrefix(args[i], "--") {
				// Break down the option into the option name and its associated value
				parts := strings.SplitN(args[i], "=", 2)
				if len(parts) != 2 {
					return "", "", "", nil, fmt.Errorf("invalid option format: %s", args[i])
				}
				name, value := parts[0], parts[1]

				// Slice of struct preserves the insertion order of the applied options
				orderedOptions = append(orderedOptions, Option{Name: name, Value: value})
			} else {
				return "", "", "", nil, fmt.Errorf("unexpected argument: %s", args[i])
			}
		}

		return command, filename, outputFilename, orderedOptions, nil
	}

	// If command is neither "header" nor "apply", then return an error
	return "", "", "", nil, fmt.Errorf("unknown command: %s", command)
}

// Reads the BMP and DIB headers from a file
func readHeaders(filename string) (*BMPHeader, *DIBHeader, error) {
	fmt.Println("Opening file: <", filename, ">")
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the BMP header info
	var bmpHeader BMPHeader
	if err := binary.Read(file, binary.LittleEndian, &bmpHeader); err != nil {
		return nil, nil, fmt.Errorf("error reading BMP header: %v", err)
	}

	if string(bmpHeader.FileType[:]) != "BM" {
		return nil, nil, errors.New("error: not a valid BMP file")
	}

	// Read the DIB header info
	var dibHeader DIBHeader
	if err := binary.Read(file, binary.LittleEndian, &dibHeader); err != nil {
		return nil, nil, fmt.Errorf("error reading DIB header: %v", err)
	}

	return &bmpHeader, &dibHeader, nil
}

// Prints the BMP and DIB header information
func printHeader(bmp *BMPHeader, dib *DIBHeader) {
	fmt.Println("BMP Header:")
	fmt.Printf("- FileType %s\n", string(bmp.FileType[:]))
	fmt.Printf("- FileSizeInBytes %d\n", bmp.FileSize)
	fmt.Printf("- HeaderSize %d\n", bmp.OffsetData)

	fmt.Println("DIB Header:")
	fmt.Printf("- DibHeaderSize %d\n", dib.DibHeaderSize)
	fmt.Printf("- WidthInPixels %d\n", dib.Width)
	fmt.Printf("- HeightInPixels %d\n", dib.Height)
	fmt.Printf("- PixelSizeInBits %d\n", dib.BitCount)
	fmt.Printf("- ImageSizeInBytes %d\n", dib.ImageSize)
}

// Reads the pixel data from the BMP file
func readPixels(filename string, bmpHeader *BMPHeader, dibHeader *DIBHeader) ([]Pixel, error) {
	return nil, nil
}

// Writes the modified pixel data to an output BMP file
func writePixels(filename string, bmpHeader *BMPHeader, dibHeader *DIBHeader, pixels []Pixel) error {
	return nil
}

// Applies horizontal or vertical mirroring
func applyMirror(pixels []Pixel, width, height int, mode string) []Pixel {
	return pixels
}

// Applies various filters like blue, red, green, grayscale, negative, pixelate or blur
func applyFilter(pixels []Pixel, width, height int, filterType string) []Pixel {
	return pixels
}

// Rotates the image by 90, 180 or 270 degrees both clockwise and counterclockwise
func applyRotate(pixels []Pixel, width, height int, angle int) []Pixel {
	return pixels
}

// Crops the image based on the given parameters
func applyCrop(pixels []Pixel, width, height, offsetX, offsetY, cropWidth, cropHeight int) []Pixel {
	return pixels
}

// Displays general usage instructions
func displayGeneralHelp() {
	fmt.Println("Usage:")
	fmt.Println("  bitmap <command> [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println("  header    prints bitmap file header information")
	fmt.Println("  apply     applies processing to the image and saves it to the file")
}

// Displays usage instructions for header command
func displayHeaderHelp() {
	fmt.Println("Usage:")
	fmt.Println("  bitmap header <source_file>")
	fmt.Println()
	fmt.Println("Description:")
	fmt.Println("  Prints bitmap file header information")
}

// Displays usage instructions for apply command
func displayApplyHelp() {
	fmt.Println("Usage:")
	fmt.Println("  bitmap apply [options] <source_file> <output_file>")
	fmt.Println()
	fmt.Println("The options are:")
	fmt.Println("  -h, --help                                                      prints program usage information")
	fmt.Println("  --mirror=<horizontal|vertical>                                  mirrors the image along the specified axis")
	fmt.Println("  --filter=<blue|red|green|grayscale|negative|pixelate|blur>      applies a specified filter to the image")
	fmt.Println("  --rotate=<right|left|90|-90|180|-180|270|-270>                  rotates the image by the specified angle")
	fmt.Println("  --crop=<offsetX-offsetY-width-height>                           crops the image based on the specified offset and dimensions")
	fmt.Println()
	fmt.Println("Note:")
	fmt.Println("  Multiple options can be combined and applied sequentially")
}

func main() {
	if len(os.Args) < 2 {
		displayGeneralHelp()
		os.Exit(1)
	}

	command, filename, outputFilename, orderedOptions, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bmpHeader, dibHeader, err := readHeaders(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	switch command {
	case "header":
		printHeader(bmpHeader, dibHeader)

	case "apply":
		pixels, err := readPixels(filename, bmpHeader, dibHeader)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Process options sequentially
		for _, opt := range orderedOptions {
			switch opt.Name {
			case "--mirror":
				pixels = applyMirror(pixels, int(dibHeader.Width), int(dibHeader.Height), opt.Value)
			case "--filter":
				pixels = applyFilter(pixels, int(dibHeader.Width), int(dibHeader.Height), opt.Value)
			case "--rotate":
				angle := 0 // Parse rotation value
				pixels = applyRotate(pixels, int(dibHeader.Width), int(dibHeader.Height), angle)
			case "--crop":
				// Parse crop values and apply cropping
				pixels = applyCrop(pixels, int(dibHeader.Width), int(dibHeader.Height), 0, 0, 100, 100)
			}
		}

		err = writePixels(outputFilename, bmpHeader, dibHeader, pixels)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

	default:
		displayGeneralHelp()
		os.Exit(1)
	}
}
