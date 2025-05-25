const apiBase = 'http://localhost:3000/api/items';

// Get item ID from URL query parameter
const urlParams = new URLSearchParams(window.location.search);
const itemId = urlParams.get('id');

if (!itemId) {
    document.getElementById('message').textContent = 'Error: No item ID provided.';
    document.getElementById('editItemForm').style.display = 'none';
} else {
    // Fetch item data and populate form
    fetch(`${apiBase}/${itemId}`)
        .then(res => {
            if (!res.ok) throw new Error('Item not found');
            return res.json();
        })
        .then(item => {
            document.getElementById('name').value = item.name;
            document.getElementById('description').value = item.description;
            document.getElementById('quantity').value = item.quantity;
            document.getElementById('location').value = item.location;
        })
        .catch(err => {
            document.getElementById('message').textContent = 'Error: ' + err.message;
            document.getElementById('editItemForm').style.display = 'none';
        });
}

// Handle form submission to update item
document.getElementById('editItemForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const updatedItem = {
        name: document.getElementById('name').value,
        description: document.getElementById('description').value,
        quantity: parseInt(document.getElementById('quantity').value),
        location: document.getElementById('location').value
    };
    try {
        const res = await fetch(`${apiBase}/${itemId}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(updatedItem)
        });
        const data = await res.json();
        if (res.ok) {
            document.getElementById('message').textContent = 'Item updated successfully!';
        } else {
            document.getElementById('message').textContent = 'Error: ' + (data.error || 'Failed to update item');
        }
    } catch (err) {
        document.getElementById('message').textContent = 'Error: ' + err.message;
    }
});
