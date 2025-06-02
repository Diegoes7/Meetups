// import { formatTimestamp } from "./utils"; //! is notwork nee to be absolute path
import { formatTimestamp } from '/templates/static/utils.js';

const DELETE_MESSAGE_MUTATION = `
  mutation DeleteMessage($messageId: ID!) {
    deleteMessage(messageID: $messageId)
  }
`;

const EDIT_MESSAGE_MUTATION = `
  mutation EditMessage($input: UpdateMessageInput!) {
    editMessage(input: $input) {
      id
      content
    }
  }
`;

export async function handleEdit(messageId) {
	const messageEl = document.getElementById(`message-content-${messageId}`);
	if (!messageEl) return;

	const originalContent = messageEl.textContent;
	const newContent = prompt('Enter new message: ', originalContent);
	if (!newContent) return;

	try {
		const res = await fetch('/query', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				query: EDIT_MESSAGE_MUTATION,
				variables: {
					input: {
						messageID: messageId,
						content: newContent,
					},
				},
			}),
		});

		const result = await res.json();
		const updatedMessage = result.data?.editMessage;
		if (updatedMessage) {
			const formatted = formatTimestamp(new Date());
			const el = document.getElementById(`message-${messageId}`);
			const htmlContent = `<div style='display: flex; align-items: center;'>
		  					<div><strong>You:</strong> <span id="message-content-${updatedMessage.id}">${updatedMessage.content}<span></div>
		  						<button class="delete-btn" data-id="${updatedMessage.id}" title="Delete" style="background:none;border:none;cursor:pointer; margin:0; padding:0.5em">🗑️</button>
		  						<button class="edit-btn" data-id="${updatedMessage.id}" title="Edit" style="background:none;border:none;cursor:pointer; margin:0; padding:0.5em">✏️</button>
								</div>
									<div>
									<span> sent: <span>${formatted}</span>
								<div>
		  				</div>`;
			if (el) el.innerHTML = htmlContent;
		} else {
			alert('Failed to edit message');
		}
	} catch (err) {
		console.error('Edit error:', err);
		alert('Something went wrong');
	}
}

export async function handleDelete(messageId) {
	const confirmed = confirm('Are you sure you want to delete this message?');
	if (!confirmed) return;

	try {
		const res = await fetch('/query', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				// Add auth headers if needed
			},
			body: JSON.stringify({
				query: DELETE_MESSAGE_MUTATION,
				variables: { messageId },
			}),
		});

		const result = await res.json();
		if (result.data?.deleteMessage) {
			// Remove from DOM or refresh messages
			document.getElementById(`message-${messageId}`)?.remove();
			console.log(document.getElementById(`message-${messageId}`));
		} else {
			alert('Failed to delete message');
		}
	} catch (err) {
		console.error('Delete error:', err);
		alert('Something went wrong');
	}
}

document.body.addEventListener('click', (e) => {
	if (e.target.closest('.delete-btn')) {
		const messageId = e.target.dataset.id;
		handleDelete(messageId);
	}
});

document.body.addEventListener('click', (e) => {
	if (e.target.closest('.edit-btn')) {
		const messageId = e.target.dataset.id;
		handleEdit(messageId);
	}
});
