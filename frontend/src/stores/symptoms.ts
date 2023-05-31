import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export type SearchSymptomsModel = {
	data: string[];
	filters: string[];
	filtered: string[];
	search: string;
};

export const createSearchSymptomsStore = (data: string[]) => {
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
