#include "embedder_c.h"
#include "embedder.hpp"

struct embedder_t_ {
  Embedder *self;
};

embedder_t *NewEmbedder(const char *path) {
    auto self = new Embedder(path);
    auto embedder = new embedder_t{self};
    return embedder;
}

void DestroyEmbedder(embedder_t *embedder) {
    delete embedder->self;
    delete embedder;
}

void Transform(embedder_t *embedder, float *image, float *prediction) {
    embedder->self->Transform(image, prediction);
}

int GetEmbeddingSize(embedder_t *embedder) {
    return embedder->self->GetEmbeddingSize();
}

int GetInputWidth(embedder_t *embedder) {
    return embedder->self->GetInputWidth();
}

int GetInputHeight(embedder_t *embedder) {
    return embedder->self->GetInputHeight();
}