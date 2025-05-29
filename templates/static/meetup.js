export async function leaveMeetup(meetupId) {
	const query = `
		mutation LeaveMeetup($meetupID: ID!) {
			leaveMeetup(meetupID: $meetupID)
		}
	`;

	try {
		const response = await fetch('/query', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				query,
				variables: {
					meetupID: meetupId,
				},
			}),
		});

		const result = await response.json();

		if (result.errors) {
			throw new Error(result.errors[0].message || 'Failed to leave meetup');
		}

		console.log('Successfully left meetup');
		// Optionally update UI
	} catch (error) {
		console.error('GraphQL leave error:', error);
		alert(error.message);
	}
}

export async function fetchInvitedUsers(meetupId) {
	const query = `
		query GetInvitedUsers($meetupID: ID!) {
			getMeetupUsersInvited(meetupID: $meetupID) {
				id
				username
				email
				firstName
				lastName
			}
		}
	`;

	try {
		const response = await fetch('/query', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				query,
				variables: {
					meetupID: meetupId,
				},
			}),
		});

		const result = await response.json();

		if (result.errors) {
			throw new Error(
				result.errors[0].message || 'Failed to fetch invited users'
			);
		}

		return result.data.getMeetupUsersInvited;
	} catch (error) {
		console.error('GraphQL invited users error:', error);
		alert(error.message);
		return [];
	}
}

export async function getParticipableMeetups() {
	const query = `
		query {
			participableMeetups {
				id
				title
				date
			}
		}
	`;

	try {
		const response = await fetch('/query', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				// Include auth token if needed
				// 'Authorization': 'Bearer YOUR_TOKEN'
			},
			body: JSON.stringify({ query }),
		});

		const result = await response.json();

		if (result.errors) {
			throw new Error(result.errors[0].message || 'Failed to fetch meetups');
		}

		console.log(
			'Meetups user can participate in:',
			result.data.participableMeetups
		);
		return result.data.participableMeetups;
	} catch (error) {
		console.error('GraphQL error:', error);
		alert(error.message);
		return [];
	}
}
