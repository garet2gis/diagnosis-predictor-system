from fastapi import FastAPI
from app.model.text_model import text_model
from pydantic import BaseModel

app = FastAPI()


class Text(BaseModel):
    text: str


@app.post("/predict")
def read_root(req: Text):
    return text_model.model.run_model(req.text)
