<script lang="ts">
	import { createSearchSymptomsStore, searchHandler } from '../stores/symptoms.ts';

	import { onDestroy } from 'svelte';

	const apiBaseUrl = import.meta.env.API_BASE_URL || 'http://0.0.0.0:2228';

	const symptoms: string[] = [
		'Кожная сыпь',
		'Дрожь',
		'Обнаружение мочеиспускания',
		'Зуд',
		'Узловые высыпания на коже'
	];

	const searchStore = createSearchSymptomsStore(symptoms);

	let copySelectedSymptoms: string[] = [];
	const unsubscribe = searchStore.subscribe((value) => {
		searchHandler(value);
		copySelectedSymptoms = value.data;
	});

	onDestroy(() => {
		unsubscribe();
	});

	type DeseaseResponse = {
		Disease: string;
		Description: string;
		Precaution: string[];
	};

	let diseaseResult: DeseaseResponse | undefined;

	const toggleSelection = (symptom) => {
		diseaseResult = undefined;

		if (copySelectedSymptoms.includes(symptom)) {
			searchStore.delete(symptom);
			return;
		}
		searchStore.push(symptom);
	};

	const handleSubmit = async () => {
		try {
			const response = await fetch(`${apiBaseUrl}/api/v1/basic_model`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(copySelectedSymptoms)
			});

			diseaseResult = await response.json();
		} catch (error) {
			alert('Произошла ошибка при отправке данных на сервер.');
		}
	};

	$: filteredList = copySelectedSymptoms.sort();
</script>

<div class="container">
	<form on:submit|preventDefault={handleSubmit} class="form">
		<div class="input">
			<div class="input-wrapper">
				<input type="text" placeholder="Поиск..." bind:value={$searchStore.search} />
			</div>

			<div class="custom-select">
				<select multiple id="symptoms">
					{#each $searchStore.filtered as symptom}
						<option
							value={symptom}
							on:click={() => toggleSelection(symptom)}
							class:symptom_selected={$searchStore.data.includes(symptom)}
						>
							{symptom}
						</option>
					{/each}
				</select>
			</div>
		</div>

		<button class="button" type="submit">Получить <br />диагноз</button>
	</form>

	<div class="symptoms_result">
		<ul class="symptoms">
			<b>Выбранные симптомы:</b>
			{#each filteredList as symptom, i}
				<li class="symptom" value={i}>{symptom}</li>
			{/each}
		</ul>

		{#if diseaseResult}
			<div class="disease_info">
				<div class="disease">
					{diseaseResult.Disease}
				</div>
				<div class="description">
					<b>Описание</b>
					<br />
					{diseaseResult.Description}
				</div>
				<ul class="precaution">
					<b>Рекомендации</b>

					<div class="recs">
						{#each diseaseResult.Precaution as p, i}
							<li value={i}>{p}</li>
						{/each}
					</div>
				</ul>
			</div>
		{:else}
			<div class="disease_info skeleton">...ваше заболевание</div>
		{/if}
	</div>
</div>

<style>
	.form {
		display: flex;
		align-items: center;
		margin: 0 36px 0 36px;
		justify-content: space-between;
	}

	input {
		padding: 0;
		min-width: 100%;
		box-sizing: border-box;
	}

	select {
		min-width: 100%;
		height: 150px;
		margin-top: 10px;
	}

	button {
		margin-top: 10px;
	}

	.symptom_selected {
		background-color: lightblue;
	}

	.skeleton {
		color: #d6d6d6;
	}

	.disease_info {
		background-color: #f5f5f5;
		padding: 10px 20px 10px 20px;
		border-radius: 5px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		font-family: Arial, sans-serif;

		flex-basis: 70%;
	}

	.description,
	.disease,
	.precaution {
		padding-bottom: 16px;
		font-size: 16px;
		line-height: 1.5;
		color: #333;
	}

	.disease {
		font-size: 42px;
	}

	.disease_info {
		margin: 20px;
	}

	ul {
		margin: 0;
		padding: 0;
	}

	.recs {
		padding-left: 10px;
	}

	.symptoms_result {
		margin-left: 16px;
		margin-right: 10px;
		display: flex;
		flex-grow: 1;
	}

	.symptoms {
		flex-basis: 30%;
		background-color: #f5f5f5;
		padding: 10px 20px 10px 20px;
		border-radius: 5px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		font-family: Arial, sans-serif;

		margin: 20px;
		font-size: 16px;
	}

	.symptom {
		padding-top: 12px;
		padding-left: 8px;
	}

	.container {
		display: flex;
		flex-direction: column;
		min-height: 90vh;
	}

	.button {
		display: inline-block;
		padding: 10px 20px;
		font-size: 18px;
		font-weight: bold;
		text-transform: uppercase;
		background-color: #1e88e5;
		color: #fff;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: background-color 0.3s ease;
	}

	.button:hover {
		background-color: #1565c0;
	}

	.input {
		flex-basis: 85%;
	}

	.input-wrapper {
		min-width: 100%;
		margin-top: 10px;

		display: flex;
		align-items: center;
	}

	.input-wrapper input {
		padding: 10px;
		font-size: 16px;
		border: 2px solid #1e88e5;
		border-radius: 4px;
		outline: none;
		transition: border-color 0.3s ease;
	}

	.input-wrapper input:focus {
		border-color: #1565c0;
	}

	.custom-select {
		min-width: 100%;
		position: relative;
		width: 200px;
		font-size: 16px;
	}

	.custom-select select {
		width: 100%;
		font-size: 16px;
		border: 2px solid #1e88e5;
		border-radius: 4px;
		appearance: none;
		outline: none;
		cursor: pointer;
	}

	.custom-select option {
		padding: 6px 10px;
	}

	.custom-select select:hover {
		border-color: #1565c0;
	}

	.custom-select select:focus {
		border-color: #1565c0;
		box-shadow: 0 0 0 2px rgba(30, 136, 229, 0.3);
	}

	.custom-select::after {
		content: '\25BC';
		position: absolute;
		top: 50%;
		right: 10px;
		transform: translateY(-50%);
		font-size: 12px;
		color: #1e88e5;
		pointer-events: none;
	}
</style>
