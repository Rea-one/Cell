import os
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation

# 获取当前脚本的绝对路径
current_script_path = os.path.abspath(__file__)
# 获取当前脚本所在目录
current_directory = os.path.dirname(current_script_path)
# 修改工作目录到脚本所在目录
os.chdir(current_directory)

DATA_DIR = 'data'  # 默认当前目录，如果你把CSV放在 data/ 下，请改为 'data'

# 加载所有CSV文件
def load_all_data():
    all_data = {}
    for filename in os.listdir(DATA_DIR):
        if filename.endswith('.csv'):
            state_id = int(filename.split('.')[0])
            df = pd.read_csv(os.path.join(DATA_DIR, filename))
            grid_data = df[[col for col in df.columns if col.startswith('格子')]].values
            all_data[state_id] = grid_data.reshape(-1, 3, 3)  # [时刻数, 3, 3]
    return all_data

all_data = load_all_data()

# 创建三维可视化动画
fig, ax = plt.subplots()
im = ax.imshow(np.zeros((3, 3)), cmap='binary', vmin=0, vmax=1)
ax.axis('off')
title_text = ax.set_title("")

# 控制当前查看的状态ID
current_state = 0

def update(frame):
    global current_state
    frames = all_data.get(current_state, [])
    if frame < len(frames):
        im.set_array(frames[frame])
        title_text.set_text(f"状态 {current_state} - 时刻 {frame}")
    return im,

ani = animation.FuncAnimation(fig, update, frames=100, interval=300, repeat=False)

# 键盘事件切换状态
def on_key(event):
    global current_state
    if event.key == 'right':
        current_state = min(max(all_data.keys()), current_state + 1)
    elif event.key == 'left':
        current_state = max(min(all_data.keys()), current_state - 1)
    print(f"当前显示状态: {current_state}")

fig.canvas.mpl_connect('key_press_event', on_key)

plt.show()