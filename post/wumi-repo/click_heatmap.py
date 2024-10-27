import json
import os
import time

import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns


"""mysql
-- 查询各链接广告位的点击坐标
select offer_id,
       ad_unit,
       max(ad_width)                                                           ad_width,
       max(ad_height)                                                          ad_height,
       max(ad_area)                                                            ad_area,
       json_arrayagg(json_array(round(click_rel_x, 2), round(click_rel_y, 2))) coords
from (select offer_id,
             ad_unit,
             click_rel_x,
             click_rel_y,
             round(ad_width)                                ad_width,
             round(ad_height)                               ad_height,
             concat(round(ad_width), 'x', round(ad_height)) ad_res,
             round(ad_width * ad_height, 2)                 ad_area
      from r_ad_click_points
      where ad_width > 0
        and ad_height > 0) b
group by offer_id, ad_unit, ad_res
having json_length(coords) > 100
order by offer_id, ad_unit, ad_area desc;
"""

def show_click_heatmap(coords, ad_width, ad_height, title="heatmap"):
    # pixel to inch
    plt.figure(figsize=(ad_width / 96, ad_height / 96))
    plt.title(title)

    x_edges = np.arange(0, ad_width, step=3)
    y_edges = np.arange(0, ad_height, step=3)
    x = [x[0] for _, x in enumerate(coords)]
    y = [y[1] for _, y in enumerate(coords)]
    heatmap, x_edge, y_edge = np.histogram2d(x, y, bins=(x_edges, y_edges))

    sns.set_style("whitegrid")
    sns.heatmap(heatmap.T, cmap="YlGnBu")

    plt.show()


only_show_offers = ['seekeptgame.com']

if __name__ == '__main__':
    path = os.getcwd() + '/click_heatmap/data_source.json'
    with open(path, 'r') as f:
        data = json.load(f)
        ad_map = {}
        for ad_coords in data:
            offer_id = ad_coords['offer_id']
            ad_unit = ad_coords['ad_unit']
            key = (offer_id, ad_unit)
            if len(only_show_offers) > 0 and offer_id not in only_show_offers:
                continue

            if key not in ad_map:
                ad_map[key] = []
            ad_map[key].append(ad_coords)

        new_ad_map = {}
        for key, ad_coords in ad_map.items():
            new_key = f"{key[0]}@{key[1]}"
            if new_key not in new_ad_map:
                new_ad_map[new_key] = {}
            if len(ad_coords) > 1:
                # scale to fit max ad area, sort descending by sql
                max_res = ad_coords[0]
                max_width, max_height = max_res["ad_width"], max_res["ad_height"]
                # modify all resolution to fit
                for i, ad_coord in enumerate(ad_coords):
                    if i == 0:
                        continue
                    scale = max(max_width / max_res["ad_width"], max_height / max_res["ad_height"])
                    for coord in ad_coord["coords"]:
                        coord[0] *= scale
                        coord[1] *= scale
                        max_res["coords"].append(coord)

            new_ad_map[new_key] = {
                "width": ad_coords[0]["ad_width"],
                "height": ad_coords[0]["ad_height"],
                "coords": ad_coords[0]["coords"]
            }

    for key, coord in new_ad_map.items():
        offer_id, ad_unit = key.split("@")
        print(coord["width"], coord["height"], f"{offer_id} - {ad_unit}")
        show_click_heatmap(coord["coords"], coord["width"], coord["height"], f"{offer_id} - {ad_unit}")
        time.sleep(1)
        # break

    # with open(os.getcwd()+"/click_heatmap/merged.json", 'w') as out_file:
    #     out_file.write(json.dumps(new_ad_map, indent=2).__str__())
