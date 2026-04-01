chromedp UI自动化

1. chromedp是go控制的chrome浏览器工具
2. 用途
    1. 爬虫
    2. 自动化操作
    3. 自动化测试
    4. 页面截图
    5. 反爬虫绕过 真人浏览器行为
3. 浏览器设置
    1. Append 在默认设置后面，加新的配置
    2. Chromed.DefaultExecAllocatorOptions[:] 把默认数据复制成一个新的slice
4. 启动Chrome进程：chromed.NewExecAllocator(context.Background (), top…)
    1. Context.background() 作用一个 根上下文  作用：控制生命周期
5. 基本使用流程
    1. 创建上下文  ctx, cancel := chromed.NewContext(context.Background())      Defer cancel
    2. 执行操作
       chromed.Run(cox, …)
    3. 写任务
       chromed.Navigate(“”)
       chromed.Click(`#btn`),
       chromed.endKeys(`#input`,”hello”),
       chromed.SendKeys(`input`,”hello”),
       4.获取结果
       var html string
       chromed.OuterHTML(“html”,&html)
       4.常用操作API
       打开页面：chromedp.Navigate(url)
       点击：	   chromed.Click(`#btn`)
       输入：        chromed.Sendkys(`#input`,”内容”)
       等待元素：chromed.WaitVisible(`#id`)
       获取文本： var text string		 chromed.Text(`#id`, &text)
       选择JS：   var res string 		chromed.Evaluate(`document.title`, &res)
       截图：	   var buf []tyte		chromed.FullScreenshot(&buf, 90)

浏览器
1. chromedp.NewExecAllocator  1.启动浏览器 2.设置浏览器参数（是否无并没有，窗口，用户数据目录，代理）
2. chromedp.NewContext 创建浏览器会话
   chromedp	现实
   ExecAllocator	打开一个 Chrome
   Context	打开一个标签页
   Action	在页面操作（点/输）
3. Chromedp.Nacigate(url) 让浏览器跳转到指定网页
4. WaitVisible(`input[placeholder="Say hi..."]`) 隐性等待，等待某个元素出现
5. Rand.Intn()生成随机数
6.

main流程图：
1. 初始化随机数
2. 准备多个profile账号
3. For循环启动coroutine
4. run User
1. 启动浏览器 exceAllocator
2. 打开页面
3. 登录
4. 进入房间
5. 无限循环：随机消息，随机时间，发送消息
1. 伪代码
    1.  配置浏览器 默认DefaultExecAllca+ 单独配置Flag（）无头 false,账号隔离 user-data-dir
    2. 启动浏览器 NewExecAllocator
    3. 在浏览器上打开标签 NewContext()
    4. Run(配置，URL，等待)
* go runUsert(profile, &wg)
1. Go是启动并发线程的关键字， &wg是指针 俄共计数并发任务


func runUserT(profile string, wg *sync.WaitGroup) {
defer wg.Done()
1. Profile string  调用者传入的内容
2. Sync.WaitGroup 指针，用来在并发场景下同步多个goroutine的完成状态
3. Wg.Done() waitGroup的计数器 1
