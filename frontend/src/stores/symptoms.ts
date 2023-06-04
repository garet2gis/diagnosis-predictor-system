import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export type SearchSymptomsModel = {
	data: string[];
	filters: string[];
	filtered: string[];
	search: string;
};

export const createSearchSymptomsStore = () => {
	const data = Object.keys(defaultSymptoms);

	const symptoms: Writable<SearchSymptomsModel> = writable({
		data: [],
		filtered: data.sort(),
		filters: data.sort(),
		search: ''
	});

	return {
		subscribe: symptoms.subscribe,
		set: symptoms.set,
		update: symptoms.update,
		push: (newSymptom: string) => {
			symptoms.update((state) => {
				state.data.push(newSymptom);
				return state;
			});
		},
		delete: (symptom: string) => {
			symptoms.update((state) => {
				state.data = state.data.filter((e) => e !== symptom);
				return state;
			});
		}
	};
};

export const searchHandler = (store: SearchSymptomsModel) => {
	const searchTerm = store.search.toLowerCase() || '';

	store.filtered = store.filters.filter((symptom) => symptom.toLowerCase().includes(searchTerm));
};

const defaultSymptoms = {
	'Боль в животе': 'abdominal_pain',
	'Ненормальная менструация': 'abnormal_menstruation',
	Кислотность: 'acidity',
	'Острая печеночная недостаточность': 'acute_liver_failure',
	'Измененный сенсориум': 'altered_sensorium',
	Беспокойство: 'restlessness',
	'Боль в спине': 'back_pain',
	'Боль в брюшной полости': 'belly_pain',
	'Черные точки': 'blackheads',
	'Дискомфорт мочевого пузыря': 'bladder_discomfort',
	Волдырь: 'blister',
	'Кровь в мокроте': 'blood_in_sputum',
	'Кровавый стул': 'bloody_stool',
	'Размытое и искаженное зрение': 'blurred_and_distorted_vision',
	Одышка: 'breathlessness',
	'Ломкие ногти': 'brittle_nails',
	Синяки: 'bruising',
	'Жжение при мочеиспускании': 'burning_micturition',
	'Боль в груди': 'chest_pain',
	Озноб: 'chills',
	'Холодные руки и ноги': 'cold_hands_and_feets',
	Кома: 'coma',
	'Заложенность носа': 'congestion',
	Запор: 'constipation',
	'Постоянное ощущение мочи': 'continuous_feel_of_urine',
	'Непрерывное чихание': 'continuous_sneezing',
	Кашель: 'cough',
	Судороги: 'cramps',
	'Темная моча': 'dark_urine',
	Обезвоживание: 'dehydration',
	Депрессия: 'depression',
	Диарея: 'diarrhoea',
	'Дисхромные пятна': 'dischromic_patches',
	'Растяжение живота': 'distention_of_abdomen',
	Головокружение: 'dizziness',
	'Сохнущие и покалывающие губы': 'drying_and_tingling_lips',
	'Увеличенная щитовидная железа': 'enlarged_thyroid',
	'Чрезмерный голод': 'excessive_hunger',
	'Внебрачные связи': 'extra_marital_contacts',
	Родословная: 'family_history',
	'Быстрый сердечный ритм': 'fast_heart_rate',
	Усталость: 'fatigue',
	'Перегрузка жидкостью': 'fluid_overload',
	'Неприятный запах мочи': 'foul_smell_ofurine',
	'Головная боль': 'headache',
	'Высокая температура': 'high_fever',
	'Боль в тазобедренном суставе': 'hip_joint_pain',
	'Давнее потребление алкоголя': 'history_of_alcohol_consumption',
	'Повышенный аппетит': 'increased_appetite',
	Несварение: 'indigestion',
	'Воспалительные ногти': 'inflammatory_nails',
	'Внутренний зуд': 'internal_itching',
	'Нерегулярный уровень сахара': 'irregular_sugar_level',
	Раздражительность: 'irritability',
	'Раздражение в анусе': 'irritation_in_anus',
	Зуд: 'itching',
	'Суставная боль': 'joint_pain',
	'Боль в колене': 'knee_pain',
	'Нехватка концентрации': 'lack_of_concentration',
	Летаргия: 'lethargy',
	'Потеря аппетита': 'loss_of_appetite',
	'Потеря баланса': 'loss_of_balance',
	'Потеря запаха': 'loss_of_smell',
	Недомогание: 'malaise',
	'Слабая лихорадка': 'mild_fever',
	'Перепады настроения': 'mood_swings',
	'Скованность движений': 'movement_stiffness',
	'Слизистая мокрота': 'mucoid_sputum',
	'Боли в мышцах': 'muscle_pain',
	'Атрофия мышц': 'muscle_wasting',
	'Мышечная слабость': 'muscle_weakness',
	Тошнота: 'nausea',
	'Боль в шее': 'neck_pain',
	'Узловые высыпания на коже': 'nodal_skin_eruptions',
	Ожирение: 'obesity',
	'Боль за глазами': 'pain_behind_the_eyes',
	'Боль во время дефекации': 'pain_during_bowel_movements',
	'Боль в анальной области': 'pain_in_anal_region',
	'Болезненная ходьба': 'painful_walking',
	'Учащенное сердцебиение': 'palpitations',
	'Проход газов': 'passage_of_gases',
	'Пятна в горле': 'patches_in_throat',
	Мокрота: 'phlegm',
	Полиурия: 'polyuria',
	'Выступающие вены на икрах': 'prominent_veins_on_calf',
	'Одутловатое лицо и глаза': 'puffy_face_and_eyes',
	'Гнойные прыщи': 'pus_filled_pimples',
	'Получение переливания крови': 'receiving_blood_transfusion',
	'Получение нестерильных инъекций': 'receiving_unsterile_injections',
	'Красная язва вокруг носа': 'red_sore_around_nose',
	'Красные пятна поверх тела': 'red_spots_over_body',
	'Покраснение глаз': 'redness_of_eyes',
	Насморк: 'runny_nose',
	'Ржавая мокрота': 'rusty_sputum',
	Суетливость: 'scurring',
	Дрожь: 'shivering',
	Аргироз: 'silver_like_dusting',
	'Давление пазухи': 'sinus_pressure',
	'Шелушение кожи': 'skin_peeling',
	'Кожная сыпь': 'skin_rash',
	'Невнятная речь': 'slurred_speech',
	'Небольшие вмятины на ногтях': 'small_dents_in_nails',
	'Вращательные движения': 'spinning_movements',
	'Обнаружение мочеиспускания': 'spotting_urination',
	'Скованность мышц шеи': 'stiff_neck',
	'Желудочное кровотечение': 'stomach_bleeding',
	'Боль в желудке': 'stomach_pain',
	'Запавшие глаза': 'sunken_eyes',
	Потливость: 'sweating',
	'Опухшие лимфатические узлы': 'swelled_lymph_nodes',
	'Отек суставов': 'swelling_joints',
	'Опухание желудка': 'swelling_of_stomach',
	'Опухшие кровяные сосуды': 'swollen_blood_vessels',
	'Опухшие конечности': 'swollen_extremeties',
	'Опухшие ноги': 'swollen_legs',
	'Раздражение горла': 'throat_irritation',
	'Ядовитый взгляд (тиф)': 'toxic_look_(typhos)',
	'Язвы на языке': 'ulcers_on_tongue',
	Неустойчивость: 'unsteadiness',
	'Визуальные нарушения': 'visual_disturbances',
	Рвота: 'vomiting',
	Слезотечение: 'watering_from_eyes',
	'Слабость в конечностях': 'weakness_in_limbs',
	'Слабость одной стороны тела': 'weakness_of_one_body_side',
	'Увеличение веса': 'weight_gain',
	'Потеря веса': 'weight_loss',
	'Желтая корочка': 'yellow_crust_ooze',
	'Желтая моча': 'yellow_urine',
	'Пожелтение глаз': 'yellowing_of_eyes',
	'Желтоватая кожа': 'yellowish_skin'
};
