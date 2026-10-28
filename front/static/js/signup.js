(function () {
  function showSuccess(form) {
    var box = document.createElement('div');
    box.className = 'form-success';
    box.setAttribute('role', 'status');

    var check = document.createElement('span');
    check.className = 'form-success-check';
    check.setAttribute('aria-hidden', 'true');
    check.textContent = '\u2713';

    var text = document.createElement('p');
    text.textContent = 'Спасибо! Форма отправлена. Наш администратор свяжется с вами в ближайшее время.';

    box.appendChild(check);
    box.appendChild(text);
    form.parentNode.replaceChild(box, form);
  }

  document.addEventListener('submit', function (e) {
    var form = e.target.closest('.signup-form');
    if (!form) return;
    e.preventDefault();

    var btn = form.querySelector('.js-signup-btn');
    if (btn) {
      btn.classList.add('is-loading');
      btn.disabled = true;
    }

    fetch(form.action, {
      method: 'POST',
      body: new FormData(form)
    })
      .then(function (res) {
        if (!res.ok) throw new Error('Request failed');
        showSuccess(form);
      })
      .catch(function () {
        if (btn) {
          btn.classList.remove('is-loading');
          btn.disabled = false;
        }
      });
  });
})();