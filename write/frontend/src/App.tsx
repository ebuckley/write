import './App.css'

function App() {
  const updateValue = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    const value = event.target.value
    console.log(value)
  }
  return (
    <div className="flex h-screen flex-col bg-white dark:bg-gray-800">
      <div className="flex justify-end p-4">
        <div className="flex gap-2">
          <button
            className="mx-2 inline-flex h-10 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium transition-colors ring-offset-background border-input bg-background hover:bg-accent hover:text-accent-foreground focus-visible:ring-ring focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50">New
            File
          </button>
          <button
            className="mx-2 inline-flex h-10 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium transition-colors ring-offset-background border-input bg-background hover:bg-accent hover:text-accent-foreground focus-visible:ring-ring focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50">Open
          </button>
          <button
            className="mx-2 inline-flex h-10 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium transition-colors ring-offset-background border-input bg-background hover:bg-accent hover:text-accent-foreground focus-visible:ring-ring focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50">Save
          </button>
          <button
            className="mx-2 inline-flex h-10 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium transition-colors ring-offset-background border-input bg-background hover:bg-accent hover:text-accent-foreground focus-visible:ring-ring focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50">Export
          </button>
        </div>
      </div>
      <div className="flex flex-grow">
        <div className="w-1/2">
          <textarea
            onChange={updateValue}
            className="flex w-full rounded-md border bg-white px-3 py-2 text-sm min-h-[80px] border-input ring-offset-background placeholder:text-muted-foreground focus-visible:ring-ring focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-gray-700"
            id="markdown-input" placeholder="Type your markdown here..."></textarea></div>
        <div className="w-1/2">
          <div className="w-full overflow-y-scroll bg-white prose dark:prose-dark dark:bg-gray-700"></div>
        </div>
      </div>
    </div>
  )
}

export default App
