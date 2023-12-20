var slideIndex = 0;
showSlides();
function showSlides() {
    var i;
    var slides = document.getElementsByClassName("mySlides");
    var dots = document.getElementsByClassName("dot");
    for (i = 0; i < slides.length; i++) {
        slides[i].style.display = "none";
    }
    slideIndex++;
    if (slideIndex > slides.length) { slideIndex = 1 }
    for (i = 0; i < dots.length; i++) {
        dots[i].className = dots[i].className.replace(" active", "");
    }
    slides[slideIndex - 1].style.display = "block";
    dots[slideIndex - 1].className += " active";
    setTimeout(showSlides, 4000);
}

$(document).ready(function () {
    var itemsPerPage = 36;
    var currentPage = 1;

    var products = $('.Tovar');
    var totalPages = Math.ceil(products.length / itemsPerPage);

    showPage(currentPage);

    function showPage(page) {
        currentPage = page;
        $('.Tovar').hide();
        var startIndex = (page - 1) * itemsPerPage;
        var endIndex = startIndex + itemsPerPage;
        products.slice(startIndex, endIndex).show();
    }

    $('.next').on('click', function () {
        if (currentPage < totalPages) {
            showPage(currentPage + 1);
        }
    });

    $('.prev').on('click', function () {
        if (currentPage > 1) {
            showPage(currentPage - 1);
        }
    });

    $('.first').on('click', function () {
        showPage(1);
    });

    $('.last').on('click', function () {
        showPage(totalPages);
    });

    updatePaginationText();

    function updatePaginationText() {
        $('.first').text('1');
        $('.last').text(totalPages);
    }
});

document.addEventListener('DOMContentLoaded', function () {
    document.querySelector('.loginForm').addEventListener('submit', function (event) {
        var email = document.getElementById('email').value.trim(); // Забираємо зайві пробіли
        var password = document.getElementById('password').value.trim(); // Забираємо зайві пробіли
        var userType = document.querySelector('input[name="userType"]:checked');
        var errorMessage = document.querySelector('.Error1'); // Отримуємо елемент p для відображення помилки

        // Виконуємо перевірку наявності значень
        if (!email || !password || !userType) {
            errorMessage.style.visibility = 'visible'; // Показуємо повідомлення про помилку
            event.preventDefault(); // Зупиняємо відправку форми
        } else {
            errorMessage.style.visibility = 'hidden'; // Ховаємо повідомлення про помилку, якщо значення введено коректно
        }
    });
});