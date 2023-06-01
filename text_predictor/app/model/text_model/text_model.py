import pandas as pd
import spacy
import pickle


class ObservationsParser:
    def __init__(self):
        self.model = spacy.load('ru_core_news_lg')
        self.di = {}

    def print_tokens(self):

        for token in self.doc:
            token_ind = token.i
            token_text = token.text
            token_pos = token.pos_
            token_dep = token.dep_
            token_head = token.head.text
            print(f"{token_ind:<10}{token_text:<12}{token_pos:<10}" \
                  f"{token_dep:<10}{token_head:<12}")

    def find_features(self):

        cur_result = {'Название признака': [], 'Значение': []}

        for token in self.doc:
            if token.i != len(self.doc) - 1:
                if self.doc[token.i].pos_ in ['NOUN', 'PROPN']:
                    cur_feature = []
                    cur_value = []
                    ind1 = 0

                    if self.doc[token.i + 1].pos_ == 'NUM':
                        ind2 = 0
                        while self.doc[token.i + ind1].pos_ != 'PUNCT':
                            cur_feature.append(self.doc[token.i + ind1].text)
                            ind1 -= 1

                        cur_feature.reverse()
                        cur_value = self.doc[token.i + 1].text
                        cur_result['Название признака'].append(' '.join(cur_feature))
                        cur_result['Значение'].append(cur_value)

                    elif self.doc[token.i + 1].text == ':':
                        ind2 = 2
                        while self.doc[token.i + ind1].pos_ != 'PUNCT':
                            cur_feature.append(self.doc[token.i + ind1].text)
                            ind1 -= 1

                        cur_feature.reverse()

                        while self.doc[token.i + ind2].pos_ != 'PUNCT':
                            cur_value.append(self.doc[token.i + ind2].text)
                            ind2 += 1

                        cur_result['Название признака'].append(' '.join(cur_feature))
                        cur_result['Значение'].append(' '.join(cur_value))

        self.di = cur_result

    def synon(self):

        proc1 = ['женский', 'женщина', 'женск', 'жен']
        proc2 = ['мужской', 'мужчина', 'мужск', 'муж']
        proc3 = ['присутствует', 'присут', 'прис', 'есть']
        proc4 = ['отсутствует', 'отсут', 'отс', 'нету', 'нет']
        res = []

        for val in self.di['Значение']:
            for p, z in zip([proc1, proc2, proc3, proc4], ['1', '0', '1', '0']):
                for i in p:
                    val = val.lower().replace(i, z)
            for p, z in zip([['ж'], ['м']], ['1', '0']):
                for i in p:
                    if len(val) == 1:
                        val = val.lower().replace(i, z)
            res.append(val.strip())
        self.di['Значение'] = res
        res = []

        proc5 = ['креатининфосфокиназа', 'креатинфосфокиназа', 'креатинкиназа']

        for val in self.di['Название признака']:
            for p, z in zip([proc5], ['кфк']):
                for i in p:
                    val = val.lower().replace(i, z)
            res.append(val.strip())
        self.di['Название признака'] = res

    def get_features(self, observation: str) -> pd.DataFrame:
        self.data = '. ' + observation + ' .'
        self.doc = self.model(self.data)
        self.find_features()
        self.synon()
        return pd.DataFrame(self.di)


class ModelHD:
    """
    модель на сердечную недостаточность

    Args:
        sample (str): строка с наблюдением
        parser (экземпляр ObservationsParser): парсер

    Returns:
        str: строка с диагнозом
    """

    def __init__(self, parser):
        self.parser = parser
        self.req_features_values_ru_1 = {'возраст': 60.0, 'анемия': 0, 'кфк': 250.0, 'фракция выброса': 38.0,
                                         'высокое кровяное давление': 0, 'тромбоциты': 262000.0, 'креатинин': 1.1,
                                         'натрий': 137.0, 'пол': 1, 'курение': 0, 'время наблюдения': 115.0}
        self.req_features_ru_1 = list(self.req_features_values_ru_1.keys())
        self.model = pickle.load(open(f'app/model/text_model/model2_rf.pkl', 'rb'))

    def feature_append(self, feat, val):
        if feat in self.df['Название признака'].values:
            self.df.loc[self.df.loc[self.df['Название признака'] == feat].index] = [feat, val]
        else:
            self.df.loc[len(self.df.index)] = [feat, val]

    def values_to_int(self):
        self.df = self.df.astype({'Значение': float})

    def requirements(self, inplace=False):
        cur_features = self.df[self.df['Название признака'].isin(self.req_features_ru_1)]['Название признака'].values
        cur_values = self.df[self.df['Название признака'].isin(self.req_features_ru_1)]['Значение'].values
        diff = list(set(self.req_features_ru_1) - set(cur_features))

        if diff:
            print(f'Отсутствующие признаки {len(diff)}/{len(self.req_features_ru_1)}:\n{diff}')
            if inplace:
                print('Автодополнение')
                for d in diff:
                    self.feature_append(d, self.req_features_values_ru_1[d])
        else:
            print('Норм')

    def load_heart_data(self, data: list):

        df = pd.DataFrame([data], columns=['age', 'anaemia', 'creatinine_phosphokinase', 'ejection_fraction',
                                           'high_blood_pressure',
                                           'platelets', 'serum_creatinine', 'serum_sodium', 'sex', 'smoking', 'time'])

        # min_max_scaler = load('enc/scaler_hd.joblib')
        min_max_scaler = pickle.load(open(f'app/model/text_model/enc/scaler_hd.pkl', 'rb'))
        # ordinal_encoder = load('enc/encoder_hd.joblib')
        ordinal_encoder = pickle.load(open(f'app/model/text_model/enc/encoder_hd.pkl', 'rb'))

        bins = [0, 50, 60, 70, 80, 90, 100]
        labels = ['<50', '50-60', '60-70', '70-80', '80-90', '90+']
        df['AgeGroup'] = pd.cut(df['age'], bins=bins, labels=labels, right=False)

        categorical_features = ['AgeGroup']
        df[categorical_features] = ordinal_encoder.transform(df[categorical_features])

        df.drop('age', axis=1, inplace=True)

        numeric_features = ['creatinine_phosphokinase', 'ejection_fraction', 'platelets', 'serum_creatinine',
                            'serum_sodium', 'time']
        df[numeric_features] = min_max_scaler.transform(df[numeric_features])

        return df

    def check_valid_data(self):
        for i in self.req_features_values_ru_1.keys():
            j = self.df.loc[self.df['Название признака'] == i]['Значение'].values[0]
            try:
                float(j)
            except ValueError:
                print(f'Неподходящий тип фичи: {i}')
                print('Замена')
                self.feature_append(i, self.req_features_values_ru_1[i])

    def run_model(self, sample):
        """
        прогноз модели

        Returns:
            str: строка с диагнозом
        """

        self.sample = sample
        self.df = self.parser.get_features(self.sample)
        df_checkpoint = self.df.copy()
        self.requirements(inplace=True)
        self.check_valid_data()
        X = []
        for i in self.req_features_values_ru_1.keys():
            X.append(self.df.loc[self.df['Название признака'] == i]['Значение'].values[0])
        X = list(map(float, X))
        X = self.load_heart_data(X)

        preds = self.model.predict_proba(X)
        preds = round(preds[0][1], 3) * 100

        res_df_checkpoint = {i[0]: i[1] for i in df_checkpoint.to_dict('split')['data']}
        res_df_checkpoint['предсказание'] = preds

        return res_df_checkpoint


parser = ObservationsParser()
model = ModelHD(parser)
