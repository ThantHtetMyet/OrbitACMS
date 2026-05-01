import './App.css'
import { Navigate, Route, Routes } from 'react-router-dom'
import SignInPage from './pages/sign-in/SignInPage'
import SignUpPage from './pages/sign-up/SignUpPage'
import UniversalTheme from './components/universal-theme/UniversalTheme'

function App() {
  return (
    <UniversalTheme>
      <main className="app-shell">
        <Routes>
          <Route path="/" element={<Navigate to="/sign-in" replace />} />
          <Route path="/login" element={<Navigate to="/sign-in" replace />} />
          <Route path="/sign-in" element={<SignInPage />} />
          <Route path="/sign-up" element={<SignUpPage />} />
        </Routes>
      </main>
    </UniversalTheme>
  )
}

export default App
