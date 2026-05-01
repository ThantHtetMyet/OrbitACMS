import { useEffect, useState } from 'react'
import dayIcon from '../../assets/icons/Day.png'
import nightIcon from '../../assets/icons/Night.png'
import './UniversalTheme.css'

function UniversalTheme({ children }) {
  const [themeMode, setThemeMode] = useState(() => {
    const storedMode = localStorage.getItem('theme-mode')
    return storedMode === 'night' ? 'night' : 'day'
  })

  useEffect(() => {
    localStorage.setItem('theme-mode', themeMode)
  }, [themeMode])

  const toggleThemeMode = () => {
    setThemeMode((currentMode) => (currentMode === 'day' ? 'night' : 'day'))
  }

  const isNightMode = themeMode === 'night'

  return (
    <div className={`universal-theme ${isNightMode ? 'theme-night' : 'theme-day'}`}>
      {children}
      <button
        className="theme-toggle"
        type="button"
        onClick={toggleThemeMode}
        aria-label={`Switch to ${isNightMode ? 'day' : 'night'} mode`}
      >
        <span className="theme-toggle__icon" aria-hidden="true">
          <img
            className="theme-toggle__icon-image"
            src={isNightMode ? dayIcon : nightIcon}
            alt=""
          />
        </span>
      </button>
    </div>
  )
}

export default UniversalTheme
