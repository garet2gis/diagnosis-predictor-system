from fastapi import FastAPI
from typing import List
from app.model.basic_model import basic_model

app = FastAPI()


@app.post("/predict")
def read_root(array_symptoms: List[str]):
    symptoms = basic_model.symps_to_labels(array_symptoms)
    pred = basic_model.model.prediction(symptoms)

    return pred
