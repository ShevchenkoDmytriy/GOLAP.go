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





