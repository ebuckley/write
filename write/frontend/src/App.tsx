import './App.css'
import {Save, SendMarkdownToRenderer, Open, Export} from "../wailsjs/go/main/App";
import {useRef, useState} from "react";

function App() {
  const [filename, setFileName] = useState('');
  const [md, setMd] = useState('');
  const ref = useRef<HTMLTextAreaElement>(null);
  const updateValue = async (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    const value = event.target.value
    const html = await SendMarkdownToRenderer(value)
    setMd(html);
  }
  const newAction = async () => {
    if (md.length > 0) {
      if (!window.confirm('Are you sure you want to create a new document?')) {
        return;
      }
    }
    setMd('');
    setFileName('');
    if (ref.current) {
     ref.current.value = '';
    }
  }
  const saveAction = async () => {
    try {
      const res = await Save()
      alert('saved!' + res);
    } catch(e) {
      alert('error saving: ' + e);
    }

  }
  const openAction = async () => {
    const res = await Open();
    setMd(res.html ?? "")
    if (ref.current) {
      ref.current.value = res.content ?? "";
    }
    setFileName(res.name ?? "")
  }
  const exportAction = async () => {
    try {
      await Export();
      alert('exported! to file');
    } catch (e) {
      // @ts-ignore shhh!
      alert("error: " + e.toString())
    }
  }
  return (
    <div className="flex h-screen flex-col bg-white dark:bg-gray-800">
      <div className="flex justify-end p-4">
        <h1>{filename}</h1>
        <div className="flex gap-2">
          <button
            onClick={newAction}
            className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 mx-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              className=" text-gray-500 dark:text-gray-200"
            >
              <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"></path>
              <polyline points="14 2 14 8 20 8"></polyline>
              <line x1="12" x2="12" y1="18" y2="12"></line>
              <line x1="9" x2="15" y1="15" y2="15"></line>
            </svg>
          </button>
          <button
            onClick={openAction}
            className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 mx-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              className=" text-gray-500 dark:text-gray-200"
            >
              <path d="m6 14 1.45-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.55 6a2 2 0 0 1-1.94 1.5H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H18a2 2 0 0 1 2 2v2"></path>
            </svg>
          </button>
          <button
            onClick={saveAction}
            className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 mx-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              className=" text-gray-500 dark:text-gray-200"
            >
              <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path>
              <polyline points="17 21 17 13 7 13 7 21"></polyline>
              <polyline points="7 3 7 8 15 8"></polyline>
            </svg>
          </button>
          <button
            onClick={exportAction}
            className="inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 mx-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
              className=" text-gray-500 dark:text-gray-200"
            >
              <path d="M12 3v12"></path>
              <path d="m8 11 4 4 4-4"></path>
              <path d="M8 5H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-4"></path>
            </svg>
          </button>
        </div>
      </div>
      <div className="flex flex-grow overflow-x-hidden">
        <div className="w-1/2 p-4 overflow-y-auto">
          <textarea
            onChange={updateValue}
            ref={ref}
            className="flex min-h-[80px] rounded-md border border-input px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 w-full h-full bg-white dark:bg-gray-700"
            id="markdown-input" placeholder="Type your markdown here..."></textarea></div>
        <div className="w-1/2 p-4 overflow-y-auto">
          <div className="w-full bg-white prose dark:prose-dark dark:bg-gray-700 markdown-content" dangerouslySetInnerHTML={{__html: md}}>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
