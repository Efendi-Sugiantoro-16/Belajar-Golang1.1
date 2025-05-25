const apiBase = 'http://localhost:3000/api/items';

async function fetchItems() {
    const res = await fetch(apiBase);
    const items = await res.json();
    const tbody = document.querySelector('#itemsTable tbody');
    tbody.innerHTML = '';
    items.forEach(item => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${item.name}</td>
            <td>${item.description}</td>
            <td>${item.quantity}</td>
            <td>${item.location}</td>
            <td>
                <button class="delete" onclick="deleteItem('${item.id}')">Delete</button>
                <button onclick="editItem('${item.id}')">Edit</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

async function deleteItem(id) {
    if (!confirm('Are you sure you want to delete this item?')) return;
    await fetch(`${apiBase}/${id}`, { method: 'DELETE' });
    fetchItems();
}

function editItem(id) {
    window.location.href = `/edit_item.html?id=${id}`;
}

fetchItems();
