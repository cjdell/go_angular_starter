window.addEventListener('load', imageGallery);

function imageGallery() {
  [].forEach.call(document.querySelectorAll('.image-gallery'), (container) => {

    [].forEach.call(container.querySelectorAll('.thumbs img'), (img) => {
      img.addEventListener('mouseover', (e) => {
        var handle = img.attributes['data-handle'].value;
        showLargeImage(handle);
      });
    });

    function showLargeImage(handle) {
      var imgs = container.querySelectorAll('.main img');
      var img = container.querySelector('.main img[data-handle=' + handle + ']');

      [].forEach.call(imgs, (img) => img.style.opacity = 0);
      img.style.opacity = 1;
    }

  });
}