<script>
const serverSrc = 'http://localhost:8080/api/v1'
const codefortag = '/codefortag'
const tags = '/tags'
const codeall = '/codeall'
const codeforid = '/codeforid'

const token = '12345678'
const hasEnter = false
const userSite = '中国钓鱼岛'
export default
{
  userSite, // 用户地址
  token, // 用户token身份
  serverSrc, // 服务器地址
  codefortag, // 通过TagID获取代码列表
  tags, // 获取全部标签
  codeall, // 获取全部代码列表
  codeforid, // 通过ID获取代码详细内容
  hasEnter // 用户登录状态
}
</script>
