// static route

import { getMeetupByID } from '/templates/static/meetup.js';

// Enables Enter key to trigger the button click
export function enableEnterKeyAction(inputSelector, buttonSelector) {
	document.addEventListener('DOMContentLoaded', () => {
		const input = document.querySelector(inputSelector);
		const button = document.querySelector(buttonSelector);

		if (!input || !button) {
			console.warn(
				`Input or button not found: ${inputSelector}, ${buttonSelector}`
			);
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

export async function fetchInvitationsByStatus(userID, status) {
	const res = await fetch('/query', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			query: `
        query GetInvitations($userID: ID, $status: InvitationStatus) {
          invitations(filter: { userID: $userID, status: $status }) {
            ID
            Status
            MeetupID
          }
        }
      `,
			variables: {
				userID,
				status,
			},
		}),
	});

	const result = await res.json();
	const invitations = result?.data?.invitations || [];

	const invitationList = document.getElementById('invitationList');

	if (!userID) {
		invitationList.innerHTML = `<p style='margin: 0.25em;'>Authenticated User is required. Please, Login.</p>`;
		return;
	}

	const meetupInvitations = invitations.map(async (invitation, inx) => {
		const m = await await getMeetupByID(invitation.MeetupID);
		const li = document.createElement('li');
		const acceptBtn =
			invitation.Status === 'pending'
				? `<button
				data-invitation-id='${invitation.ID}'
				data-meetup-name='${m.name}'
				class='accept_btn'
			>
				Accept
			</button>`
				: '';
		const declineBtn =
			invitation.Status === 'pending'
				? `<button
				data-invitation-id='${invitation.ID}'
				data-meetup-name='${m.name}'
				class='decline_btn'
			>
				Decline
			</button>`
				: '';
		li.innerHTML = `<div style="display:flex; justify-content: space-between; align-items: baseline; 
		border: 1px solid #ccc; border-radius: .25em; padding: .5em; margin-bottom: 0.5em; gap:2em">
			<div style="display: flex; gap: 0.5em; justify-content: flex-start; align-items: baseline;">
			${inx + 1}. Invitation ID: ${invitation.ID} — Meetup: ${m.name} 
			</div>
						<div style="display:flex; flex-wrap: wrap; gap: 0.5em" >
							${acceptBtn} 
							${declineBtn} 
						</div>
		</div>`;
		invitationList.appendChild(li);
	});

	invitationList.innerHTML = '';

	if (invitations.length === 0) {
		invitationList.innerHTML = `<p style='margin: 0.25em;'>No invitations found for status "${status}".</p>`;
		return;
	}

	return meetupInvitations;
}

export async function acceptInvitation(invitationId) {
	const res = await fetch('/query', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			query: `
        mutation AcceptInvitation($invitationID: ID!) {
          acceptInvitation(invitationID: $invitationID) {
            ID
            Status
          }
        }
      `,
			variables: { invitationID: invitationId },
		}),
	});

	const result = await res.json();
	const accepted = result?.data?.acceptInvitation;

	if (!accepted) {
		console.error('Failed to accept invitation:', result.errors);
	} else {
		console.log(
			`Accepted invitation ${accepted.ID}, new status: ${accepted.Status}`
		);
	}
}

export async function declineInvitation(invitationId) {
	const res = await fetch('/query', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({
			query: `
        mutation DeclineInvitation($invitationID: ID!) {
          declineInvitation(invitationID: $invitationID) {
            ID
            Status
          }
        }
      `,
			variables: { invitationID: invitationId },
		}),
	});

	const result = await res.json();
	const declined = result?.data?.declineInvitation;

	if (!declined) {
		console.error('Failed to decline invitation:', result.errors);
	} else {
		console.log(
			`Declined invitation ${declined.ID}, new status: ${declined.Status}`
		);
	}
}

export const BASE_URL = window.location.hostname.includes('localhost')
	? 'http://localhost:8080'
	: 'https://meetups-y7ke.onrender.com';
