#include "embedder.hpp"

Embedder::Embedder(std::string model_path) : model_path_(model_path) {
  module_ = torch::jit::load(model_path);
}
void Embedder::Transform(float* image, float* prediction) const {
  // TODO: run image on the model and output results by the prediction ptr
}
int Embedder::GetEmbeddingSize() const {
  // TODO: get this from model congig
  return 0;
}
int Embedder::GetInputWidth() const {
  // TODO: get this from model congig
  return 0;
}
int Embedder::GetInputHeight() const {
  // TODO: get this from model congig
  return 0;
}
