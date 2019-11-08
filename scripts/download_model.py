import torch
import torchvision
import os
from torch import nn
import json

# https://github.com/pytorch/pytorch/issues/2486
class Lin_View(nn.Module):
	def __init__(self):
		super(Lin_View, self).__init__()
	def forward(self, x):
		return x.view(x.size()[0], -1)

def build_model():
    model = torchvision.models.resnet50(pretrained=True)
    embedder = nn.Sequential(
        *list(model.children())[:-2], 
        nn.MaxPool2d(7), 
        Lin_View()
    )
    for param in embedder.parameters():
        param.requires_grad = False
    return embedder

def main():
    print("Start building model")
    model = build_model()
    print("Complete building")
    print(model)
    iw, ih = 224, 224
    example = torch.rand(1, 3, iw, ih)
    out = model(example)
    print(f"Output shape = {out.shape}")
    assert len(out.shape) == 2, "Output should be 2 dims"

    config = {
        "input_width": iw,
        "input_height": ih,
        "output_shape": out.shape[1]
    }

    print("Start tracing")
    traced_script_module = torch.jit.trace(model, example)
    print("Complete tracing")

    path = "model.pt"
    traced_script_module.save(path)
    json.dump(config, open("config.json", "w"))
    print(f"Model is saved at {os.path.abspath(path)}")


if __name__ == "__main__":
    main()