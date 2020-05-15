#include <grpc++/grpc++.h>
#include <spdlog/spdlog.h>

#include <memory>

#include "embedder_server.h"

int main() {
  grpc::ServerBuilder builder;
  auto address = "0.0.0.0:1234";
  builder.AddListeningPort(address, grpc::InsecureServerCredentials());
  EmbedderServer embedder_server{};
  builder.RegisterService(&embedder_server);
  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
  spdlog::info("Server run on {0} address", address);
  server->Wait();
  return 0;
}
