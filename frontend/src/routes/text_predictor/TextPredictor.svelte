<script>
	let inputText = '';
	let probability = null;
	let data = null;

	const apiBaseUrl = import.meta.env.API_BASE_URL || 'http://0.0.0.0:2228';

	async function handleSubmit() {
		if (inputText !== '') {
			try {
				const response = await fetch(`${apiBaseUrl}/api/v1/text_model`, {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({ text: inputText })
				});

				data = await response.json();

				probability = data['предсказание'];

				delete data['предсказание'];
			} catch (error) {
				console.error('Ошибка при отправке запроса:', error);
			}
		}
	}
</script>

<div class="content">
	<div class="input">
		<div class="text">
			<textarea
				class="textarea"
				id="inputText"
				bind:value={inputText}
				placeholder="Введите данные пациента"
			/>
		</div>

		<button
			class="button"
			on:click={handleSubmit}
			disabled={inputText.trim() === ''}
			class:disabled={inputText.trim() === ''}
			>Отправить
		</button>
	</div>
	{#if probability !== null}
		<p>Риск летального исхода от сердечной недостаточности: {probability}%</p>

		<b>Выделенные признаки из текста</b>
		<table>
			<thead>
				<tr>
					<th>Поле</th>
					<th>Значение</th>
				</tr>
			</thead>
			<tbody>
				{#each Object.entries(data) as [key, value]}
					{#if typeof value === 'string'}
						<tr>
							<td>{key}</td>
							<td>{value}</td>
						</tr>
					{/if}
				{/each}
			</tbody>
		</table>
	{/if}
</div>

<style>
	.content {
		margin: 10px 36px;
	}

	textarea {
		width: 100%;
		min-height: 300px;
		margin-right: 30px;
		padding: 10px;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	button {
		padding: 10px 20px;
		font-size: 16px;
		background-color: #1e88e5;
		color: #fff;
		border: none;
		border-radius: 4px;
		cursor: pointer;
	}

	p {
		margin-top: 20px;
	}

	.button {
		max-height: 60px;
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

	.textarea {
		padding: 10px;
		font-size: 20px;
		border: 2px solid #1e88e5;
		border-radius: 4px;
		resize: vertical;
		outline: none;
		transition: border-color 0.3s ease;
	}

	.textarea:hover,
	.textarea:focus {
		border-color: #1e88e5;
	}

	table {
		width: 100%;
		border-collapse: collapse;
		margin-top: 20px;
	}

	th,
	td {
		padding: 12px;
		text-align: left;
		border-bottom: 1px solid #ddd;
	}

	th {
		font-weight: bold;
		background-color: #f5f5f5;
	}

	td {
		background-color: #fff;
	}

	tr:nth-child(even) td {
		background-color: #f9f9f9;
	}

	tr:hover td {
		background-color: #f1f1f1;
	}

	.disabled {
		background-color: gray;
		cursor: not-allowed;
	}

	.input {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
	}

	.text {
		flex-basis: 80%;
	}
</style>
