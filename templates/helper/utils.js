// static/js/util.js

// Enables Enter key to trigger the button click
export function enableEnterKeyAction(inputSelector, buttonSelector) {
	document.addEventListener('DOMContentLoaded', () => {
		const input = document.querySelector(inputSelector);
		const button = document.querySelector(buttonSelector);

		if (!input || !button) {
			console.warn(`Input or button not found: ${inputSelector}, ${buttonSelector}`);
			return;
		}

		input.addEventListener('keydown', (event) => {
			if (event.key === 'Enter') {
				event.preventDefault();
				button.click();
			}
		});
	});
}

// Format ISO string to human-readable format: hh:mm:ss dd/mm/yyyy
export function formatTimestamp(isoString) {
	const date = new Date(isoString);
	const hours = date.getHours().toString().padStart(2, '0');
	const minutes = date.getMinutes().toString().padStart(2, '0');
	const seconds = date.getSeconds().toString().padStart(2, '0');
	const day = date.getDate().toString().padStart(2, '0');
	const month = (date.getMonth() + 1).toString().padStart(2, '0');
	const year = date.getFullYear();

	return `${hours}:${minutes}:${seconds} ${day}/${month}/${year}`;
}
