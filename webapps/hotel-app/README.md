# hotel-app 酒店预定用户端

### 安装
- 环境 node: v16.20.2, 可以使用nvm工具管理node版本，非常方便。

- 安装依赖并运行: npm i && npm run dev

### Zustand 笔记

初次使用 Zustand React 状态管理库，写写笔记。



set()  get() 
immer(set, get) => ({})

const useXXXStore = create<Type>()(
    immer((set, get) => ({
        xxx: xxx
        func() {}
    }))
)

提出问题，假如两个state有x、y个状态，同时对应两个X、Y组件分别单独管理，
当你使用set更新x状态时，X组件重新渲染是没问题的，但是Y组件也重新渲染了，
这个就比较耗资源了，因此就有selector。

更新X、Y组件都会渲染
const { obj:{x}, changeX } = useXYStore();

// selector 语法, 上面的x状态更新，就不会影响到Y组件了
const y = useXYStore(state => state.obj.Y)