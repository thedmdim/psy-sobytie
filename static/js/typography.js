(function () {
  var short = [
    'а', 'в', 'и', 'к', 'о', 'с', 'у', 'я',
    'за', 'из', 'от', 'по', 'до', 'на', 'не', 'ни',
    'но', 'да', 'то', 'же', 'ли', 'бы', 'он', 'мы',
    'вы', 'во', 'ко', 'со', 'об', 'под', 'над',
    'без', 'при', 'для', 'или', 'что', 'как', 'так',
    'это', 'вот', 'всё', 'все', 'ещё', 'уже',
    'там', 'тут', 'где', 'кто', 'чем', 'тех'
  ];

  var escaped = short.map(function (w) {
    return w.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  });
  var pattern = new RegExp('(^|[\\s.,!?;:"()«»—–-])(' + escaped.join('|') + ')\\s+', 'gi');

  var skip = { SCRIPT: 1, STYLE: 1, CODE: 1, PRE: 1, TEXTAREA: 1, INPUT: 1, NOSCRIPT: 1 };

  var walker = document.createTreeWalker(
    document.body,
    NodeFilter.SHOW_TEXT,
    null
  );

  var node;
  while ((node = walker.nextNode())) {
    if (skip[node.parentNode.nodeName]) continue;
    node.nodeValue = node.nodeValue.replace(pattern, function (match, before, word) {
      return before + word + '\u00A0';
    });
  }
})();