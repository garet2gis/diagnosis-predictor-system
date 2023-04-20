import json
import numpy as np
import pickle


class ModelPrediction:
    def __init__(self, model):
        self._model = model

    def __json_to_dict(self):
        with open('app/model/basic_model/json/description.json', 'r') as f:
            self.di_description = json.load(f)
        with open('app/model/basic_model/json/precaution.json', 'r') as f:
            self.di_precaution = json.load(f)
        with open('app/model/basic_model/json/disease_eng_to_ru.json', 'r') as f:
            self.disease_eng_to_ru = json.load(f)
        with open('app/model/basic_model/json/descr_eng_to_ru.json', 'r') as f:
            self.descr_eng_to_ru = json.load(f)
        with open('app/model/basic_model/json/precs_eng_to_ru.json', 'r') as f:
            self.precs_eng_to_ru = json.load(f)

    def prediction(self, data: list) -> dict:
        self.__json_to_dict()
        pred = self._model.predict([data])[0].strip()
        descr = self.di_description[pred]
        precaution = self.di_precaution[pred]
        return {'Disease': self.disease_eng_to_ru[pred],
                'Description': self.descr_eng_to_ru[descr],
                'Precaution': [self.precs_eng_to_ru[prec] for prec in precaution]}


with open('app/model/basic_model/json/severity_weights.json', 'r') as f:
    severity_weights = json.load(f)

with open('app/model/basic_model/json/symp_ru_to_eng.json', 'r') as f:
    symp_ru_to_eng = json.load(f)


def symps_to_labels(symps):
    symps = [symp_ru_to_eng[s] for s in symps]
    arr = np.array([severity_weights[symp] for symp in symps])
    return np.pad(arr, (0, 17 - len(arr))).tolist()


model = ModelPrediction(pickle.load(open('app/model/basic_model/model1_rf.pkl', 'rb')))
