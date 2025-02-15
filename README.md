# Question 1
GO evaluation for DSA course. This is the solution for Question 1 only.

# Path Converter in Go

This Go program converts a relative path into an absolute path based on the given current directory.

## Installation and Usage

Follow these steps to clone and run the project:

### 1. Clone the Repository
Ensure you have Git installed, then run:

```sh
git clone https://github.com/rouhan14/question1.git

```

### 2. Move to the question1 directory

```sh
cd question1

```

### 3. Run all the files with the following command to run the program
Ensure you have GO installed, then run:

```sh
go run .

```

### Test Cases

Run the program with the following test cases:

#### Test Case 1

Input:
```sh
currentDir := "/home/user/docs/rouhan"
relativePath := "./pictures/image.jpg"
```

Expected Output:

```sh
Absolute Path: /home/user/docs/rouhan/pictures/image.jpg
```
#### Test Case 2
Input:

```sh
currentDir := "/home/user/docs/rouhan"
relativePath := "../music/song.mp3"
```
Expected Output:

```sh
Absolute Path: /home/user/docs/music/song.mp3
```
#### Test Case 3
Input:

```sh
currentDir := "/home/user/docs/rouhan"
relativePath := "../../music/song.mp3"
```
Expected Output:

```sh
Absolute Path: /home/user/music/song.mp3
```
#### Test Case 4
Input:

```sh
currentDir := "/home/user/docs/rouhan"
relativePath := "../"
```
Expected Output:

```sh
Absolute Path: /home/user/docs
```