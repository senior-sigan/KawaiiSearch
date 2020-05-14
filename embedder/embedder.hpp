#pragma once

#include <torch/script.h>

#include <string>

class Embedder {
  std::string model_path_;
  torch::jit::script::Module module_;

 public:
  Embedder(std::string model_path);
  void Transform(float* image, float* prediction) const;
  int GetEmbeddingSize() const;
  int GetInputWidth() const;
  int GetInputHeight() const;
};
