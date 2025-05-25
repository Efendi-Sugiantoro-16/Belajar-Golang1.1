const apiBase = 'http://localhost:3000/api/items';

document.getElementById('addItemForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const item = {
        name: document.getElementById('name').value,
        description: document.getElementById('description').value,
        quantity: parseInt(document.getElementById('quantity').value),
        location: document.getElementById('location').value
    };
    try {
        const res = await fetch(apiBase, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(item)
        });
        if (res.ok) {
            document.getElementById('message').textContent = 'Item added successfully!';
            e.target.reset();
        } else {
            const error = await res.json();
            document.getElementById('message').textContent = 'Error: ' + (error.error || 'Failed to add item');
        }
    } catch (err) {
        document.getElementById('message').textContent = 'Error: ' + err.message;
    }
});
