package main

/*
typedef struct embedder_t_ embedder_t;
embedder_t *NewEmbedder(const char *path);
void DestroyEmbedder(embedder_t *embedder);
void Transform(embedder_t *embedder, float *image, float *prediction);
int GetEmbeddingSize(embedder_t *embedder);
int GetInputWidth(embedder_t *embedder);
int GetInputHeight(embedder_t *embedder);

#cgo LDFLAGS: -L${SRCDIR}/src_cpp/cmake-build-debug/ -lkawaii_search
#cgo darwin LDFLAGS:  -L${SRCDIR}/src_cpp/libs/osx/libtorch/lib   -lstdc++ -ltorch -lc10
#cgo linux LDFLAGS:   -L${SRCDIR}/src_cpp/libs/linux/libtorch/lib -lstdc++ -ltorch -lc10
*/
import "C"
import "unsafe"
import "fmt"

// Embedder encapsulates raw-c object to intercat with neural network
type Embedder struct {
	raw *C.embedder_t
}

// NewEmbedder creates Embedder object
func NewEmbedder(filename string) *Embedder {
	cModelPath := C.CString(filename)
	//TODO: path a model config to the NewEmbedder constructor
	predictor := C.NewEmbedder(cModelPath)
	return &Embedder{
		raw: predictor,
	}
}

// DestroyEmbedder frees the memory of the Embedder structure
// because it's raw-c object
func (embedder *Embedder) DestroyEmbedder() {
	C.DestroyEmbedder(embedder.raw)
}

// Transform pass the image through the neural network to get an image embedding vector
// The image should be in the [r1,r2..rN...,g1,g2..gN...,b1,b2..bN] format
// The pixel format should be with zero mean and 1-std.
func (embedder *Embedder) Transform(image []float32) []float32 {
	out := make([]float32, embedder.GetEmbeddingSize())

	cImage := (*C.float)(unsafe.Pointer(&image[0]))
	cOut := (*C.float)(unsafe.Pointer(&out[0]))

	C.Transform(embedder.raw, cImage, cOut)
	return out
}

// GetEmbeddingSize returns the neural network output size
func (embedder *Embedder) GetEmbeddingSize() int {
	return int(C.GetEmbeddingSize(embedder.raw))
}

// GetInputWidth returns the neural network input image width size
func (embedder *Embedder) GetInputWidth() int {
	return int(C.GetInputWidth(embedder.raw))
}

// GetInputHeight returns the neural network input image height size
func (embedder *Embedder) GetInputHeight() int {
	return int(C.GetInputHeight(embedder.raw))
}

func (embedder *Embedder) String() string {
	return fmt.Sprintf("Embedder{Height=%d, Width=%d, Output=%d}", embedder.GetInputHeight(), embedder.GetInputWidth(), embedder.GetEmbeddingSize())
}