# Rclone Password Reveal

A simple Go script to decrypt passwords obfuscated with `rclone obscure`.

## 📌 Features
- Decrypts passwords stored in `rclone.conf` that were obfuscated with `rclone obscure`.
- Uses the `obscure.Reveal` function from `rclone`'s internal package.
- Easy to use from the command line.

## 🛠 Installation
### Prerequisites
- Go 1.19 or later installed ([Download Go](https://go.dev/dl/))

### Clone the Repository
```sh
git clone https://github.com/yourusername/rclone-password-reveal.git
cd rclone-password-reveal
```

### Install Dependencies
Since the repository already contains `go.mod` with the correct dependencies, simply run:
```sh
go mod tidy
```

## 🚀 Usage

### Run the Script
To reveal an obfuscated password, use:
```sh
go run main.go "your_obscured_password"
```
Example:
```sh
go run main.go "jKbX8aq3wGjZorrMVcM_tQtL5BxA5cvc_DQ"
```
Output:
```
Original password: mypassword
```

### Build the Executable (Optional)
If you want to use the script as a standalone executable:
```sh
go build -o rclone_password_reveal
```
Then, run it with:
```sh
./rclone_password_reveal "jKbX8aq3wGjZorrMVcM_tQtL5BxA5cvc_DQ"
```

## 📖 How It Works
This script utilizes the `obscure.Reveal` function from `rclone`, which is part of the `github.com/rclone/rclone/fs/config/obscure` package. It takes an obfuscated password as input and reveals the original password.

## ⚠️ Disclaimer
This script is for **educational purposes only**. Use it responsibly and do not use it for unauthorized access.

## 📝 License
This project is open-source under the MIT License.

## 🙌 Contributing
Feel free to submit issues and pull requests to improve the project!

