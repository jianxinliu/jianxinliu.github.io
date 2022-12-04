import base64
import io
import sys

import pandas as pd
import numpy as np
from scipy.interpolate import griddata
import matplotlib.pyplot as plt

def contour(xarr, yarr, zarr):
    column = ['x', 'y', 'z']
    array = np.array([xarr, yarr, zarr])
    frame = pd.DataFrame(data=array, index=column).T

    df1 = frame['x']
    df2 = frame['y']
    df3 = frame['z']
    
    odf1 = np.linspace(df1.min(), df1.max(), 100)
    odf2 = np.linspace(df2.min(), df2.max(), 100)
    odf1, odf2 = np.meshgrid(odf1, odf2)
    odf3_new = griddata((df1, df2), df3, (odf1, odf2), method='cubic', fill_value=df3.min())
    
    plt.figure(figsize=(9, 6))
    
    plt.contourf(odf1, odf2, odf3_new,
                 levels=np.arange(odf3_new.min(), odf3_new.max(), (odf3_new.max() - odf3_new.min()) / 15),
                 cmap='nipy_spectral',
                 extend='both')
    
    line = plt.contour(odf1, odf2, odf3_new,
                       levels=np.arange(odf3_new.min(), odf3_new.max(), (odf3_new.max() - odf3_new.min()) / 15))
    plt.clabel(line, inline=True, fontsize=12)
    
    string_io_bytes = io.BytesIO()
    plt.savefig(string_io_bytes, format='jpg')
    string_io_bytes.seek(0)
    base64_data = base64.b64encode(string_io_bytes.read()).decode("utf-8").replace("\n", "")
    return 'data:image/jpg;base64,' + base64_data

if __name__ == '__main__':
    a = []
    for i in range(1, len(sys.argv)):
        a.append([float(x) for x in sys.argv[i].split(',')])

    # 需要将结果输出到 stdout， java 才好获取
    print(contour(a[0],a[1],a[2]))