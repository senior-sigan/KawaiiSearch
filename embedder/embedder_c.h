#pragma once

#if __cplusplus
#define C_API extern "C"
#else
#define C_API
#endif

typedef struct embedder_t_ embedder_t;

C_API
embedder_t *NewEmbedder(const char *path);

C_API
void DestroyEmbedder(embedder_t *embedder);

C_API
void Transform(embedder_t *embedder, float *image, float *prediction);

C_API
int GetEmbeddingSize(embedder_t *embedder);

C_API
int GetInputWidth(embedder_t *embedder);

C_API
int GetInputHeight(embedder_t *embedder);