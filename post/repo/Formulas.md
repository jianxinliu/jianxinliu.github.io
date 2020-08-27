# 统计学相关公式

标准差：$ \sigma=\sqrt{\frac{\Sigma_i^n(X_i - \bar{X})^2}{n-1}}$ 

方差： $var = \sigma^2$

精确度：$CP = \frac{usl - lsl}{6\sigma}$ 

过程能力指标：$CPK=min(CPU,CPL),CPU=\frac{usl-\bar{X}}{3\sigma},CPL=\frac{\bar{X}-lsl}{3\sigma}$

一组点一元回归线的斜率 $\hat{\beta_1}$：$slope = \frac{n\Sigma_{i=0}^n x_iy_i - \Sigma_{i=0}^n x_i\Sigma_{i=0}^n y_i}{n\Sigma_{i=0}^n x_i x_i - \Sigma_{i=0}^n x_i \Sigma_{i=0}^n x_i}$

一组点一元回归线的截距 $\hat{\beta_0}$：$intercept = \bar{y}-slope * \bar{x}$ 

一组点一元回归线方程：$y = \hat{\beta_0} + \hat{\beta_1} x$

正态分布累计分布函数： $F(x_i,\mu,\sigma) = \frac{1}{\sigma\sqrt{2\pi}}\int_{-\infty}^x \exp^{-\frac{(x-\mu)^2}{2\sigma^2}}dx$。在程序中求积分使用的是近似的方法，具体参考 `Udf.normalSDist`

相关性系数：$r = \frac{\Sigma_{i=0}^n (x_i - \bar{x})(y_i - \bar{y})}{\sqrt{\Sigma_{i=0}^n(x_i-\bar{x})^2 \Sigma_{i=0}^n (y_i - \bar{y})^2}}$

两列的对应值平方差的和 `sumx2my2`:$\Sigma_0^i{(arrX_i^2-arrY_i^2)}$  `sumx2my2` => $sum \space x^2 \space minus \space y^2$

两列的对应值平方和的和 `sumx2py2`: $\Sigma_0^i{(arrX_i^2+arrY_i^2)}$ `sumx2py2` => $sum \space x^2 \space plus \space y^2$

[$\Gamma$ 函数]([https://zh.wikipedia.org/wiki/%CE%93%E5%87%BD%E6%95%B0](https://zh.wikipedia.org/wiki/Γ函数))： $\Gamma(n) = (n-1)!$

[$\Beta$ 函数]([https://zh.wikipedia.org/wiki/%CE%92%E5%87%BD%E6%95%B0](https://zh.wikipedia.org/wiki/Β函数))：$\Beta(x,y)=\frac{\Gamma(x)\Gamma(y)}{\Gamma(x+y)}=\frac{(x-1)!(y-1)!}{(x+y-1)!}$

概率密度函数： `pdf`

分布函数（累计分布函数）：`cdf`

## F 分布

`pdf`:$f(x,df_1,df_2)=\frac{1}{\Beta(\frac{df_1}{2},\frac{df_2}{2})} (\frac{df_1}{df_2})^{\frac{df_1}{2}} x^{\frac{df_1}{2} - 1} (1+\frac{df_1}{df_2}x)^{-\frac{df_1+df_2}{2}} , df_1,df_2 > 2$

另一种形式(使用 $\Gamma$ 函数替换)：`pdf`:
$$
f(x,df_1,df_2)=\left\{
\begin{aligned}
\frac{\Gamma(\frac{df_1 + df_2}{2})}{\Gamma(\frac{df_1}{2})\Gamma(\frac{df_2}{2})} (\frac{df_1}{df_2})^{\frac{df_1}{2}} \frac{x^{\frac{df_1}{2}-1}} {(1+\frac{df_1}{df_2}x)^{\frac{df_1+df_2}{2}}} ,x > 0\\
0,x\le 0 
\end{aligned}
\right. \\
df_1,df_2 > 2
$$


## T分布

`pdf`: $f_T(t)=\frac{\Gamma(\frac{df+1}{2})}{\Gamma(\frac{df}{2})\sqrt{\pi df}} (1+\frac{t^2}{df})^{-\frac{df + 1}{2}}$ 

`cdf`: $F_T(t) = 1- \frac{1}{2}I_{x(t)}(\frac{df}{2},\frac{1}{2})$,$I_{x(t)}$ 为不完全$\Beta$ 函数，$x(t)=\frac{df}{df + t^2}$



## 卡方分布

`pdf`:
$$
f_k(x)=\left\{
\begin{aligned}
\frac{\frac{1}{2}^{\frac{df}{2}}}{\Gamma(\frac{df}{2})} x^{\frac{df}{2}-1} e^-{\frac{x}{2}},x > 0 \\ 
0,x \le 0
\end{aligned}
\right.
$$
`cdf` : $F_k(x)=\frac{\gamma(\frac{df}{2}),\frac{x}{2}}{\Gamma(\frac{df}{2})}$ $\gamma$ 是不完全 $\Gamma$ 函数



## logNormal 分布

`pdf`:$f(x,\mu,\sigma)=\frac{1}{x\sigma\sqrt{2\pi}}e^-{\frac{(\ln{x}-\mu)^2}{2\sigma^2}}=\frac{normal.pdf(ln{x},\mu,\sigma)}{x}$

`cdf`:$\frac{1}{2}[1+erf(\frac{\ln{x}-\mu}{\sqrt{2\sigma^2}})]$



具体实现可参考 [jStat](https://github.com/jstat/jstat)，[smile-Statistical Machine Intelligence and Learning Engine](https://github.com/haifengl/smile/tree/master/math/src/main/java/smile/stat/distribution)，原理可参考[wikipedia](https://zh.wikipedia.org/) 相关词条

