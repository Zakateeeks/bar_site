document.addEventListener('DOMContentLoaded', () => {
    fetch('/api/drinks')
        .then(response => response.json())
        .then(data => {
            const drinkList = document.getElementById('drink-list');
            data.forEach(drink => {
                const drinkCard = document.createElement('div');
                drinkCard.classList.add('drink-card');

                drinkCard.innerHTML = `
                    <img src="${drink.imageUrl}" alt="${drink.name}">
                    <div class="drink-info">
                        <h3>${drink.name}</h3>
                        <p>${drink.description}</p>
                        <p>Цена: ${drink.price}₽</p>
                    </div>
                `;

                drinkList.appendChild(drinkCard);
            });
        })
        .catch(error => console.error('Ошибка при загрузке напитков:', error));
});
