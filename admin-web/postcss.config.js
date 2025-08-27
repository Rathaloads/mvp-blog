export default {
  plugins: {
    "@tailwindcss/postcss": {},
    autoprefixer: {},
    "postcss-pxtorem": {
        rootValue: 16, // 基准字体大小
        unitPrecision: 5, // 转换后保留的小数位数
        propList: ['*'], // 需要转换的属性列表
        selectorBlackList: [], // 不需要转换的选择器
        replace: true, // 是否直接替换值
        mediaQuery: false, // 是否转换媒体查询中的值
        minPixelValue: 1, // 最小像素值
        exclude: [/node_modules/i] // 排除 node_modules
    }
  },
}