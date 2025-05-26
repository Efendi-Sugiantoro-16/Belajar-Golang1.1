const ctx = document.getElementById('chartCanvas').getContext('2d');
const chartTitle = document.getElementById('chartTitle');
const chartTypeSelector = document.getElementById('chartTypeSelector');
let chartInstance = null;
let rawData = [];

// Fetch data from API
async function fetchData() {
    try {
        const response = await fetch('/api/items');
        if (!response.ok) throw new Error('Network response was not ok');
        rawData = await response.json();
        renderChart('bar');
    } catch (error) {
        alert('Failed to load data: ' + error.message);
    }
}

// Process data for charts
function processData() {
    // Aggregate quantity by name
    const quantityByName = {};
    // Aggregate quantity by location
    const quantityByLocation = {};
    // For line chart: prepare data points per name (quantity) and average quantity
    const quantities = [];

    rawData.forEach(item => {
        // Sum quantity by name
        quantityByName[item.name] = (quantityByName[item.name] || 0) + item.quantity;
        // Sum quantity by location
        quantityByLocation[item.location] = (quantityByLocation[item.location] || 0) + item.quantity;
        quantities.push(item.quantity);
    });

    // Calculate average quantity
    const averageQuantity = quantities.length ? quantities.reduce((a,b) => a+b, 0) / quantities.length : 0;

    return { quantityByName, quantityByLocation, averageQuantity };
}

// Render chart based on type
function renderChart(type) {
    if (chartInstance) {
        chartInstance.destroy();
    }

    const { quantityByName, quantityByLocation, averageQuantity } = processData();

    if (type === 'bar') {
        chartTitle.textContent = 'Bar Chart - Total Quantity by Name';
        chartInstance = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: Object.keys(quantityByName),
                datasets: [{
                    label: 'Total Quantity',
                    data: Object.values(quantityByName),
                    backgroundColor: 'rgba(54, 162, 235, 0.7)',
                    borderColor: 'rgba(54, 162, 235, 1)',
                    borderWidth: 1
                }]
            },
            options: {
                responsive: true,
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: 'Quantity'
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: 'Item Name'
                        }
                    }
                },
                plugins: {
                    legend: { display: false },
                    tooltip: { enabled: true }
                }
            }
        });
    } else if (type === 'line') {
        chartTitle.textContent = 'Line Chart - Quantity Comparison and Average';
        const labels = Object.keys(quantityByName);
        const dataValues = Object.values(quantityByName);
        const averageLine = new Array(labels.length).fill(averageQuantity);

        chartInstance = new Chart(ctx, {
            type: 'line',
            data: {
                labels: labels,
                datasets: [
                    {
                        label: 'Quantity',
                        data: dataValues,
                        borderColor: 'rgba(75, 192, 192, 1)',
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        fill: true,
                        tension: 0.3
                    },
                    {
                        label: 'Average Quantity',
                        data: averageLine,
                        borderColor: 'rgba(255, 99, 132, 1)',
                        borderDash: [5, 5],
                        fill: false,
                        pointRadius: 0
                    }
                ]
            },
            options: {
                responsive: true,
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: 'Quantity'
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: 'Item Name'
                        }
                    }
                },
                plugins: {
                    tooltip: { enabled: true }
                }
            }
        });
    } else if (type === 'pie') {
        chartTitle.textContent = 'Pie Chart - Percentage Distribution by Name and Location';
        // Combine name and location for pie chart labels
        const labels = [];
        const dataValues = [];
        const backgroundColors = [];

        // Add name distribution
        Object.entries(quantityByName).forEach(([name, qty], idx) => {
            labels.push(`Name: ${name}`);
            dataValues.push(qty);
            backgroundColors.push(`hsl(${(idx * 50) % 360}, 70%, 60%)`);
        });
        // Add location distribution
        Object.entries(quantityByLocation).forEach(([loc, qty], idx) => {
            labels.push(`Location: ${loc}`);
            dataValues.push(qty);
            backgroundColors.push(`hsl(${(idx * 50 + 180) % 360}, 70%, 60%)`);
        });

        chartInstance = new Chart(ctx, {
            type: 'pie',
            data: {
                labels: labels,
                datasets: [{
                    data: dataValues,
                    backgroundColor: backgroundColors,
                    borderColor: '#fff',
                    borderWidth: 1
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: { position: 'right' },
                    tooltip: { enabled: true }
                }
            }
        });
    }
}

chartTypeSelector.addEventListener('change', (e) => {
    renderChart(e.target.value);
});

fetchData();
