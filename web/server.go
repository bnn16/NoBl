package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nobl/evaluator"
	"nobl/lexer"
	"nobl/object"
	"nobl/parser"
	"os"
	"sync"
	"time"
)

// REPLRequest represents the JSON request from the frontend
type REPLRequest struct {
	Code      string `json:"code"`
	SessionID string `json:"sessionId,omitempty"`
}

// REPLResponse represents the JSON response to the frontend
type REPLResponse struct {
	Result    string   `json:"result"`
	Errors    []string `json:"errors,omitempty"`
	SessionID string   `json:"sessionId"`
}

// SessionEnvironment stores the environment for a session
type SessionEnvironment struct {
	Env      *object.Environment
	LastUsed time.Time
}

// Sessions stores all active session environments
var (
	sessions = make(map[string]*SessionEnvironment)
	mu       sync.Mutex
)

// cleanupSessions removes old sessions (run this in a goroutine)
func cleanupSessions() {
	for {
		time.Sleep(10 * time.Minute)
		mu.Lock()
		now := time.Now()
		for id, session := range sessions {
			// Remove sessions not used for more than 30 minutes
			if now.Sub(session.LastUsed) > 30*time.Minute {
				delete(sessions, id)
			}
		}
		mu.Unlock()
	}
}

// handleEval handles code evaluation requests from the frontend
func handleEval(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for more secure cloud deployment
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only accept POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req REPLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Create a response
	var response REPLResponse

	// Get or create session environment
	mu.Lock()
	var sessionEnv *SessionEnvironment
	var exists bool

	if req.SessionID != "" {
		sessionEnv, exists = sessions[req.SessionID]
	}

	if !exists {
		// Create a new session ID if one doesn't exist or is invalid
		req.SessionID = fmt.Sprintf("%d", time.Now().UnixNano())
		sessionEnv = &SessionEnvironment{
			Env:      object.NewEnvironment(),
			LastUsed: time.Now(),
		}
		sessions[req.SessionID] = sessionEnv

		// Start the cleanup goroutine if this is the first session
		if len(sessions) == 1 {
			go cleanupSessions()
		}
	} else {
		// Update last used time
		sessionEnv.LastUsed = time.Now()
	}

	// Set the session ID in the response
	response.SessionID = req.SessionID

	// Get the environment for this session
	env := sessionEnv.Env
	mu.Unlock()

	// Evaluate the code
	l := lexer.New(req.Code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		response.Errors = p.Errors()
	} else {
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			response.Result = evaluated.Inspect()
		}
	}

	// Send back the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleRoot serves the main HTML page
func handleRoot(w http.ResponseWriter, r *http.Request) {
	// Check if there's a code parameter for sharing (unused for now but kept for future enhancements)
	_ = r.URL.Query().Get("code")

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nobl Programming Language</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/codemirror.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/theme/dracula.min.css">
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 1000px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f5f5f5;
            color: #333;
        }
        header {
            text-align: center;
            margin-bottom: 30px;
        }
        h1 {
            color: #2c3e50;
        }
        .logo {
            font-size: 2.5em;
            font-weight: bold;
            margin-bottom: 0;
        }
        .subtitle {
            font-style: italic;
            color: #7f8c8d;
        }
        .repl-container {
            background-color: #2c3e50;
            border-radius: 5px;
            padding: 15px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            margin-bottom: 20px;
        }
        .output-area {
            background-color: #282a36;
            border-radius: 5px;
            padding: 10px;
            height: 300px;
            overflow-y: auto;
            margin-bottom: 15px;
            font-family: monospace;
            white-space: pre-wrap;
            color: #f8f8f2;
        }
        .controls {
            display: flex;
            justify-content: space-between;
            margin-top: 10px;
        }
        .command {
            color: #8be9fd;
            margin-bottom: 5px;
        }
        .result {
            color: #50fa7b;
            margin-bottom: 10px;
        }
        .error {
            color: #ff5555;
            margin-bottom: 10px;
        }
        .CodeMirror {
            height: auto;
            font-family: 'Fira Code', 'Courier New', monospace;
            border-radius: 5px;
        }
        button {
            background-color: #bd93f9;
            color: white;
            border: none;
            border-radius: 5px;
            padding: 8px 15px;
            cursor: pointer;
            transition: background-color 0.3s;
            font-weight: bold;
        }
        button:hover {
            background-color: #9580ff;
        }
        .button-group {
            display: flex;
            gap: 10px;
        }
        .tabs {
            display: flex;
            margin-bottom: 10px;
        }
        .tab {
            padding: 8px 15px;
            cursor: pointer;
            background-color: #44475a;
            color: #f8f8f2;
            border: none;
            margin-right: 5px;
            border-radius: 5px 5px 0 0;
        }
        .tab.active {
            background-color: #6272a4;
        }
        .tab-content {
            display: none;
        }
        .tab-content.active {
            display: block;
        }
        .examples {
            margin-top: 30px;
        }
        .example-code {
            background-color: #282a36;
            border: 1px solid #44475a;
            border-radius: 5px;
            padding: 10px;
            font-family: monospace;
            white-space: pre-wrap;
            margin-bottom: 15px;
            cursor: pointer;
            transition: background-color 0.3s;
            color: #f8f8f2;
        }
        .example-code:hover {
            background-color: #44475a;
        }
        .share-url {
            padding: 8px;
            border: 1px solid #ddd;
            border-radius: 5px;
            width: 100%;
            margin-top: 10px;
            font-family: monospace;
            display: none;
        }
        footer {
            margin-top: 30px;
            text-align: center;
            color: #7f8c8d;
            font-size: 14px;
            padding-top: 20px;
            border-top: 1px solid #ddd;
        }
        .language-features {
            display: flex;
            justify-content: space-between;
            flex-wrap: wrap;
            margin: 20px 0;
        }
        .feature-card {
            background-color: #fff;
            border-radius: 5px;
            padding: 15px;
            width: 30%;
            margin-bottom: 20px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        .feature-card h3 {
            color: #2c3e50;
            margin-top: 0;
        }
        @media (max-width: 768px) {
            .feature-card {
                width: 100%;
            }
        }
    </style>
</head>
<body>
    <header>
        <div class="logo">NoBl</div>
        <h1>Programming Language</h1>
        <p class="subtitle">A modern, elegant programming language with no bloat</p>
    </header>

    <div class="tabs">
        <button class="tab active" onclick="openTab(event, 'repl')">REPL</button>
        <button class="tab" onclick="openTab(event, 'about')">About</button>
        <button class="tab" onclick="openTab(event, 'docs')">Documentation</button>
    </div>

    <div id="repl" class="tab-content active">
        <div class="repl-container">
            <div id="output" class="output-area"></div>
            <div class="editor-container">
                <textarea id="codeEditor"></textarea>
            </div>
            <div class="controls">
                <div class="button-group">
                    <button id="runButton">Run Code</button>
                    <button id="clearButton">Clear Output</button>
                </div>
                <div class="button-group">
                    <button id="shareButton">Share Code</button>
                    <button id="resetButton" title="Reset the session environment">Reset Session</button>
                </div>
            </div>
            <input type="text" id="shareUrl" class="share-url" readonly>
        </div>

        <div class="examples">
            <h2>Code Examples</h2>
            <div class="example-code" onclick="setExample(this)">let five = 5;
let ten = 10;
let add = fn(x, y) { x + y };
let result = add(five, ten);
result</div>

            <div class="example-code" onclick="setExample(this)">let max = fn(x, y) { 
    if (x > y) { 
        x 
    } else { 
        y 
    } 
};
max(5, 10)</div>

            <div class="example-code" onclick="setExample(this)">
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2)
        }
    }
};
fibonacci(10)</div>
        </div>
    </div>

    <div id="about" class="tab-content">
        <h2>About NoBl</h2>
        <p>NoBl (No Bloat) is a lightweight, efficient interpreted programming language focusing on simplicity and performance. 
        Following a functional paradigm, the project involves developing a complete language implementation, including lexer, parser, AST, and evaluator.</p>
        
        <p>The core philosophy of NoBl is to provide developers with a tool that minimizes unnecessary overhead, 
        allowing for rapid development and efficient execution.</p>

        <div class="language-features">
            <div class="feature-card">
                <h3>First-class Functions</h3>
                <p>Functions are treated as values and can be passed around, returned, and assigned to variables.</p>
            </div>
            <div class="feature-card">
                <h3>Closures</h3>
                <p>Functions can access variables from their containing scope, creating powerful closure patterns.</p>
            </div>
            <div class="feature-card">
                <h3>Simple Syntax</h3>
                <p>Clean, familiar syntax that makes code easy to read and write.</p>
            </div>
            <div class="feature-card">
                <h3>Built-in Data Structures</h3>
                <p>Support for arrays and hash maps for easy data manipulation.</p>
            </div>
            <div class="feature-card">
                <h3>Expression-based</h3>
                <p>Everything is an expression, leading to concise and expressive code.</p>
            </div>
            <div class="feature-card">
                <h3>No Runtime Dependencies</h3>
                <p>Self-contained implementation with no external dependencies.</p>
            </div>
        </div>
    </div>

    <div id="docs" class="tab-content">
        <h2>Documentation</h2>
        
        <h3>Basic Syntax</h3>
        <div class="example-code">
// Variable declaration
let name = "NoBl";
let age = 25;
let isAwesome = true;

// Function definition
let greet = fn(person) {
    "Hello, " + person + "!"
};

// Function call
greet("World");
        </div>

        <h3>Data Types</h3>
        <div class="example-code">
// Integers
let age = 25;

// Strings
let name = "NoBl";

// Booleans
let isTrue = true;
let isFalse = false;

// Arrays
let fruits = ["apple", "banana", "orange"];
fruits[0];  // "apple"

// Hash Maps
let person = {"name": "John", "age": 30};
person["name"];  // "John"
        </div>

        <h3>Control Flow</h3>
        <div class="example-code">
// If/Else expression
let result = if (10 > 5) {
    "10 is greater than 5"
} else {
    "This will never be returned"
};

// Functions with conditions
let max = fn(a, b) {
    if (a > b) {
        a
    } else {
        b
    }
};
        </div>
    </div>

    <footer>
        <p>© 2025 NoBl Programming Language</p>
        <p>Developed with ❤️ by Bogdan</p>
    </footer>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/codemirror.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.65.2/mode/javascript/javascript.min.js"></script>
    <script>
        const outputArea = document.getElementById('output');
        const shareUrlInput = document.getElementById('shareUrl');
        const runButton = document.getElementById('runButton');
        const clearButton = document.getElementById('clearButton');
        const shareButton = document.getElementById('shareButton');
        const resetButton = document.getElementById('resetButton');

        // Initialize CodeMirror for syntax highlighting
        const editor = CodeMirror.fromTextArea(document.getElementById('codeEditor'), {
            lineNumbers: true,
            mode: 'javascript',
            theme: 'dracula',
            autoCloseBrackets: true,
            matchBrackets: true,
            tabSize: 2,
            indentWithTabs: false,
        });

        // Store the session ID
        let sessionId = localStorage.getItem('noblSessionId') || '';

        // Check for shared code in URL
        const urlParams = new URLSearchParams(window.location.search);
        const sharedCode = urlParams.get('code');
        if (sharedCode) {
            editor.setValue(decodeURIComponent(sharedCode));
        }

        // Show welcome message
        outputArea.innerHTML = '<div class="result">Welcome to the NoBl Programming Language REPL!</div>';
        outputArea.innerHTML += '<div class="result">Type your code and click Run to execute it.</div>';
        outputArea.innerHTML += '<div class="result">Your session state is preserved between evaluations.</div>';

        // Function to send code to the server
        async function evaluateCode(code) {
            // Add command to output
            outputArea.innerHTML += '<div class="command">>> ' + code.replace(/</g, '&lt;').replace(/>/g, '&gt;') + '</div>';
            
            try {
                const response = await fetch('/eval', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ 
                        code: code,
                        sessionId: sessionId
                    })
                });
                
                const data = await response.json();
                
                // Save the session ID for future requests
                if (data.sessionId) {
                    sessionId = data.sessionId;
                    localStorage.setItem('noblSessionId', sessionId);
                }
                
                if (data.errors && data.errors.length > 0) {
                    outputArea.innerHTML += '<div class="error">Error: ' + data.errors.join('\nError: ').replace(/</g, '&lt;').replace(/>/g, '&gt;') + '</div>';
                } else if (data.result) {
                    outputArea.innerHTML += '<div class="result">' + data.result.replace(/</g, '&lt;').replace(/>/g, '&gt;') + '</div>';
                } else {
                    outputArea.innerHTML += '<div class="result">null</div>';
                }
            } catch (error) {
                outputArea.innerHTML += '<div class="error">Network error: ' + error.message + '</div>';
            }
            
            // Scroll to bottom
            outputArea.scrollTop = outputArea.scrollHeight;
        }

        // Run button click handler
        runButton.addEventListener('click', () => {
            const code = editor.getValue().trim();
            if (code) {
                evaluateCode(code);
            }
        });

        // Clear button click handler
        clearButton.addEventListener('click', () => {
            outputArea.innerHTML = '<div class="result">Output cleared.</div>';
        });

        // Share button click handler
        shareButton.addEventListener('click', () => {
            const code = editor.getValue().trim();
            if (code) {
                // Create a shareable URL
                const shareUrl = window.location.origin + window.location.pathname + '?code=' + encodeURIComponent(code);
                shareUrlInput.value = shareUrl;
                shareUrlInput.style.display = 'block';
                shareUrlInput.select();
                document.execCommand('copy');
                alert('Share URL copied to clipboard!');
            }
        });

        // Reset session button click handler
        resetButton.addEventListener('click', () => {
            // Clear the session ID to force creation of a new environment
            sessionId = '';
            localStorage.removeItem('noblSessionId');
            outputArea.innerHTML = '<div class="result">Session reset. All defined variables and functions have been cleared.</div>';
        });

        // Function to set example code
        function setExample(element) {
            editor.setValue(element.textContent);
            editor.focus();
        }

        // Tab navigation
        function openTab(evt, tabName) {
            // Hide all tab content
            const tabContent = document.getElementsByClassName("tab-content");
            for (let i = 0; i < tabContent.length; i++) {
                tabContent[i].className = tabContent[i].className.replace(" active", "");
            }

            // Remove "active" class from all tabs
            const tabs = document.getElementsByClassName("tab");
            for (let i = 0; i < tabs.length; i++) {
                tabs[i].className = tabs[i].className.replace(" active", "");
            }

            // Show the specific tab content
            document.getElementById(tabName).className += " active";
            
            // Add "active" class to the button that opened the tab
            evt.currentTarget.className += " active";
        }
    </script>
</body>
</html>`))
}

// handleHealth is a health check endpoint for cloud platforms
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "up",
		"version": "1.0", // You might want to read this from version.txt
	})
}

// ServeWebREPL starts the web server for the REPL
func ServeWebREPL(port int) {
	// Register handlers
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/eval", handleEval)
	http.HandleFunc("/health", handleHealth)

	// Get port from environment variable (for Heroku)
	portStr := os.Getenv("PORT")
	if portStr != "" {
		fmt.Sscanf(portStr, "%d", &port)
	}

	fmt.Printf("Starting NoBl Web REPL server on port %d...\n", port)
	fmt.Printf("Open http://localhost:%d in your browser\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
