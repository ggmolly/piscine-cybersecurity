import os
import shutil
import random
import string
import numpy
from pathlib import Path
from PIL import Image

random_string = lambda x: ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(x))
PATH = os.path.join(str(Path.home()), "infection")
IMAGE_SIZE = (64, 64)

def fill_directory(path: str):
    # Create 5 random images, and 5 random text files with random strings
    for i in range(5):
        output_path = os.path.join(path, random_string(10)) + ".jpg" + (".WCRY" if i != 0 else "")
        print("[#] Creating a random image to", output_path)
        image = Image.new("RGB", IMAGE_SIZE)
        # fill with random pixels
        pixels = numpy.array(image)
        pixels[:, :] = numpy.random.randint(0, 255, size=(64, 64, 3), dtype=numpy.uint8)
        image = Image.fromarray(pixels)
        image.save(output_path, "JPEG")
        if i == 1:
            # remove read permissions
            os.chmod(output_path, 0o222)
        if i == 2:
            # remove write permissions
            os.chmod(output_path, 0o444)


    for i in range(5):
        output_path = os.path.join(path, random_string(10)) + ".txt" + (".WCRY" if i != 0 else "")
        print("[#] Creating a random text file to", output_path)
        with open(output_path.format(os.path.join(path, random_string(10))), "w") as f:
            f.write(random_string(256))
        if i == 1:
            # remove read permissions
            os.chmod(output_path, 0o222)
        if i == 2:
            # remove write permissions
            os.chmod(output_path, 0o444)

if __name__ == "__main__":
    shutil.rmtree(PATH)
    os.mkdir(PATH)
    random_dirs = [random_string(10) for _ in range(3)]
    for d in random_dirs:
        os.mkdir(os.path.join(PATH, d))
        fill_directory(os.path.join(PATH, d))
    fill_directory(PATH)