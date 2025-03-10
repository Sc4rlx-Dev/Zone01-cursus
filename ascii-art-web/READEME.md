# ASCII Art Web Application

A simple yet elegant **Go-powered web application** that converts user input into ASCII art using various font styles. Built with a minimalist **frontend** and a robust **backend**, it features error handling, sleek UI design, and a smooth user experience.

## 🚀 Features

- 🎨 **Three ASCII Fonts:** Supports `standard`, `shadow`, and `thinkertoy` styles.
- 🌐 **Web-based UI:** Designed with **glassmorphism**, animations, and a responsive layout.
- 🛠 **Error Handling:** Custom error pages for `400`, `404`, and `500` HTTP responses.
- ⚡ **Fast Processing:** Optimized ASCII conversion with efficient string manipulation.
- 🔐 **Security Measures:** Prevents XSS attacks by escaping ASCII outputs.

---

## 📋 Prerequisites

- **Go 1.16+** installed
- A modern **web browser** (Chrome, Firefox, Edge, etc.)

---

## 📥 Installation

### **1️⃣ Clone the repository**
```bash
git clone https://github.com/yourusername/ascii-art-web.git
cd ascii-art-web
```

### **2️⃣ Install dependencies** *(No external dependencies required)*
```bash
go mod tidy
```

### **3️⃣ Run the server**
```bash
go run main.go
```
The server will start at **http://localhost:8080**.

---

## 🏗 Project Structure
```
📂 ascii-art-web/
│── 📂 static/              # Static assets (CSS, images, error pages)
│   │── styles.css          # Styles for main page
│   │── style-res.css       # Styles for result page
│   │── 400.html            # Bad Request error page
│   │── 404.html            # Not Found error page
│   │── 500.html            # Internal Server Error page
│── 📂 function/            # Helper functions for ASCII processing
│   │── helper.go           # Functions for splitting, hashing, and mapping ASCII art
│── main.go                # Entry point for the web server
│── go.mod                 # Go module file
│── README.md              # Project documentation
```

---

## 🔧 Usage

### **1️⃣ Open the Web Interface**
Navigate to:
```
http://localhost:8080
```

### **2️⃣ Input Text and Select a Font**
- Enter text in the input field.
- Choose a font (`standard`, `shadow`, or `thinkertoy`).
- Click the **Run** button.

### **3️⃣ View ASCII Art Result**
- The ASCII output is displayed in a stylish UI.
- Click **Back** to return to the input page.

---

## ⚙️ How It Works

1. **User submits text** via the form.
2. **The Go server processes input**:
   - Reads the selected font file (`.txt`).
   - Converts input into ASCII format.
   - Escapes HTML to prevent XSS attacks.
3. **ASCII output is returned** in an HTML page with styled formatting.

---

## 📜 API Endpoints

### `GET /`
- Serves the main HTML page (`index.html`).

### `POST /ascii-art`
- Accepts form data:
  - `user_input`: The text to be converted.
  - `option`: The font choice.
- Returns an HTML page with ASCII art output.

---

## 🛠 Error Handling

| Status Code | Meaning | Action Taken |
|------------|---------|--------------|
| **400** | Bad Request | Invalid or missing form data |
| **404** | Not Found | Page does not exist |
| **500** | Internal Server Error | Font file missing or processing error |

---

## 🎨 UI & Styling
- Uses **glassmorphism** for modern aesthetics.
- `@keyframes` animations for form transitions.
- Custom CSS themes for input fields and buttons.

---

## 🤝 Contributing

1. **Fork the repository**
2. **Create a new branch**
   ```bash
   git checkout -b feature-new
   ```
3. **Make changes & commit**
   ```bash
   git commit -m "Added new feature"
   ```
4. **Push to GitHub & create a PR**
   ```bash
   git push origin feature-new
   ```

---

## ⚠️ Known Issues

- **Special characters may break ASCII alignment** → Needs a validation function.
- **Performance drops with large inputs** → Possible optimization with caching.

---

## 📄 License

@Sc4rlx-Dev
[0k4run](https://github.com/0k4run)
[]morida(google.com)

---

### ✨ **Enjoy your ASCII Art Journey!** ✨

