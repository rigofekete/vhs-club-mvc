import { useRef, useEffect, useState } from 'react'

function ParabolicBackground({ animate = false, animateDuration = 0, trigger = 0 }) {
  const canvasRef = useRef(null)
  const animFrameRef = useRef(null)
  const lastTimeRef = useRef(0)
  const stopTimerRef = useRef(null)
  const [animating, setAnimating] = useState(animate || animateDuration > 0)

  // Animate on mount if animateDuration is set
  useEffect(() => {
    if (animateDuration > 0) {
      setAnimating(true)
      stopTimerRef.current = setTimeout(() => setAnimating(false), animateDuration)
      return () => clearTimeout(stopTimerRef.current)
    }
  }, [])

  // Re-trigger animation when trigger prop changes (skip initial mount)
  const isFirstMount = useRef(true)
  useEffect(() => {
    if (isFirstMount.current) {
      isFirstMount.current = false
      return
    }
    if (animateDuration > 0) {
      // Clear any existing stop timer
      if (stopTimerRef.current) clearTimeout(stopTimerRef.current)
      setAnimating(true)
      stopTimerRef.current = setTimeout(() => setAnimating(false), animateDuration)
    }
  }, [trigger])

  useEffect(() => {
    if (animate) setAnimating(true)
  }, [animate])

  function drawFrame(canvas, ctx, t) {
    const w = canvas.width
    const h = canvas.height

    ctx.clearRect(0, 0, w, h)

    // Wave back and forth using sin — never drifts away
    const wave = Math.sin(t * 0.015) * 30

    const layers = [
      { count: 12, opacity: 0.06, lineWidth: 3, offset: 0 },
      { count: 10, opacity: 0.12, lineWidth: 2, offset: 40 },
      { count: 8, opacity: 0.22, lineWidth: 1.2, offset: 80 },
      { count: 6, opacity: 0.35, lineWidth: 0.7, offset: 120 },
    ]

    // Horizontal parabolas
    layers.forEach((layer) => {
      for (let i = 0; i < layer.count; i++) {
        const p = i / (layer.count - 1)
        const centerY = h * 0.1 + p * h * 0.8
        const spread = 0.3 + (1 - layer.opacity) * 0.7
        const layerWave = wave * (layer.offset / 40 + 1)

        ctx.beginPath()
        ctx.strokeStyle = `rgba(90, 100, 200, ${layer.opacity})`
        ctx.lineWidth = layer.lineWidth

        const startX = -layer.offset + layerWave
        ctx.moveTo(startX, centerY)
        ctx.quadraticCurveTo(
          w * spread + layer.offset + layerWave * 0.5,
          centerY - h * 0.4 * (0.5 - p) + layerWave,
          w + layer.offset + layerWave * 0.3,
          centerY * 0.6
        )
        ctx.stroke()
      }
    })

    // Vertical parabolas
    layers.forEach((layer) => {
      for (let i = 0; i < layer.count; i++) {
        const p = i / (layer.count - 1)
        const centerX = w * 0.05 + p * w * 0.9
        const layerWave = wave * (layer.offset / 40 + 1)

        ctx.beginPath()
        ctx.strokeStyle = `rgba(70, 80, 180, ${layer.opacity * 0.7})`
        ctx.lineWidth = layer.lineWidth

        const startY = -layer.offset + layerWave
        ctx.moveTo(centerX, startY)
        ctx.quadraticCurveTo(
          centerX + w * 0.15 * (0.5 - p) + layerWave * 0.5,
          h * 0.5 + layer.offset + layerWave,
          centerX * 0.8 + w * 0.1,
          h + layer.offset + layerWave * 0.3
        )
        ctx.stroke()
      }
    })

    // Diagonal depth lines
    const vanishX = w * 0.5
    const vanishY = h * 0.35

    for (let i = 0; i < 16; i++) {
      const angle = (i / 16) * Math.PI * 2
      const reach = 1.2
      const endX = vanishX + Math.cos(angle) * w * reach + wave * Math.cos(angle)
      const endY = vanishY + Math.sin(angle) * h * reach + wave * Math.sin(angle)

      const depth = (i % 4) / 4
      const opacity = 0.04 + depth * 0.06

      ctx.beginPath()
      ctx.strokeStyle = `rgba(80, 90, 190, ${opacity})`
      ctx.lineWidth = 0.5 + depth * 0.5
      ctx.moveTo(vanishX, vanishY)
      ctx.quadraticCurveTo(
        (vanishX + endX) / 2 + Math.sin(angle) * 50,
        (vanishY + endY) / 2 + Math.cos(angle) * 50,
        endX,
        endY
      )
      ctx.stroke()
    }
  }

  // Initial draw + resize handler
  useEffect(() => {
    const canvas = canvasRef.current
    const ctx = canvas.getContext('2d')

    function resize() {
      canvas.width = window.innerWidth
      canvas.height = window.innerHeight
      drawFrame(canvas, ctx, lastTimeRef.current)
    }

    resize()
    window.addEventListener('resize', resize)

    return () => {
      window.removeEventListener('resize', resize)
      if (animFrameRef.current) cancelAnimationFrame(animFrameRef.current)
    }
  }, [])

  // Start/stop animation loop
  useEffect(() => {
    const canvas = canvasRef.current
    const ctx = canvas.getContext('2d')

    if (animating) {
      let time = lastTimeRef.current

      function loop() {
        time++
        lastTimeRef.current = time
        drawFrame(canvas, ctx, time)
        animFrameRef.current = requestAnimationFrame(loop)
      }

      loop()

      return () => {
        if (animFrameRef.current) cancelAnimationFrame(animFrameRef.current)
      }
    }
    // When animation stops, canvas stays as-is — no redraw, no snap
  }, [animating])

  return (
    <canvas
      ref={canvasRef}
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100vw',
        height: '100vh',
        pointerEvents: 'none',
      }}
    />
  )
}

export default ParabolicBackground
