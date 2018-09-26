module.exports = [{
  name: '基础',
  component: 'Layout/Layout',
  sub: [{
    name: '首页',
    showPath: '/',
    componentName: 'index-pages/Main',
    active: 1
  }, {
    name: '选画中心',
    showPath: 'SearchShuHua',
    componentName: 'search-pages/SearchShuHua',
    active: 2
  }, {
    name: '艺术家',
    showPath: 'SearchArtist',
    componentName: 'search-pages/searchArtist',
    active: 2
  }]
}]
