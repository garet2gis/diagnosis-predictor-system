<header>
	<h1 class="logo">МЕДПОМОЩЬ</h1>
	<div class="navigation-buttons">
		<a href="/">
			<button>Модель 1</button>
		</a>
		<a href="/text_predictor">
			<button>Модель 2</button>
		</a>
	</div>
</header>

<style>
	header {
		display: flex;
		align-items: center;

		background-color: #1e88e5;
		padding-left: 20px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
	}

	.logo {
		font-size: 32px;
		font-weight: bold;
		color: #333;
	}

	.navigation-buttons {
		padding-left: 20px;
		display: flex;
		gap: 10px;
		height: 100%;
	}

	.navigation-buttons button {
		height: 100%;

		padding: 10px 20px;
		font-size: 16px;
		color: #545454;
		background-color: #f5f5f5;
		border: none;
		border-radius: 4px;
		cursor: pointer;
		transition: background-color 0.3s ease;
	}

	.navigation-buttons button:hover {
		background-color: #1565c0;
		color: #fff;
		transition: background-color 0.5s ease;
	}
</style>
