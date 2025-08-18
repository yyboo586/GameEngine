/**
 * 浏览器指纹生成工具
 * 用于为匿名用户生成唯一标识
 */

class BrowserFingerprint {
    constructor() {
        this.fingerprint = null;
        this.sessionId = null;
    }

    /**
     * 生成浏览器指纹
     * @returns {string} 指纹哈希
     */
    generateFingerprint() {
        if (this.fingerprint) {
            return this.fingerprint;
        }

        const components = [
            this.getCanvasFingerprint(),
            this.getScreenInfo(),
            this.getTimeZone(),
            this.getLanguage(),
            this.getPlatform(),
            this.getHardwareConcurrency(),
            this.getDeviceMemory(),
            this.getConnectionInfo(),
            this.getWebGLInfo(),
            this.getFontList()
        ];

        // 组合所有组件
        const fingerprintString = components.filter(Boolean).join('|');
        
        // 生成哈希（简单实现，生产环境建议使用更安全的哈希算法）
        this.fingerprint = this.simpleHash(fingerprintString);
        
        return this.fingerprint;
    }

    /**
     * 获取Canvas指纹
     */
    getCanvasFingerprint() {
        try {
            const canvas = document.createElement('canvas');
            const ctx = canvas.getContext('2d');
            
            // 绘制一些文本和图形
            ctx.textBaseline = 'top';
            ctx.font = '14px Arial';
            ctx.fillText('Browser Fingerprint 🎯', 2, 2);
            ctx.fillStyle = 'rgba(102, 204, 0, 0.7)';
            ctx.fillRect(100, 5, 50, 20);
            
            return canvas.toDataURL();
        } catch (e) {
            return null;
        }
    }

    /**
     * 获取屏幕信息
     */
    getScreenInfo() {
        return [
            screen.width,
            screen.height,
            screen.colorDepth,
            screen.pixelDepth,
            window.devicePixelRatio
        ].join('x');
    }

    /**
     * 获取时区信息
     */
    getTimeZone() {
        return Intl.DateTimeFormat().resolvedOptions().timeZone;
    }

    /**
     * 获取语言信息
     */
    getLanguage() {
        return navigator.language || navigator.userLanguage;
    }

    /**
     * 获取平台信息
     */
    getPlatform() {
        return navigator.platform;
    }

    /**
     * 获取CPU核心数
     */
    getHardwareConcurrency() {
        return navigator.hardwareConcurrency || 'unknown';
    }

    /**
     * 获取设备内存（如果支持）
     */
    getDeviceMemory() {
        return navigator.deviceMemory || 'unknown';
    }

    /**
     * 获取网络连接信息
     */
    getConnectionInfo() {
        if (navigator.connection) {
            return [
                navigator.connection.effectiveType,
                navigator.connection.rtt,
                navigator.connection.downlink
            ].join('-');
        }
        return 'unknown';
    }

    /**
     * 获取WebGL信息
     */
    getWebGLInfo() {
        try {
            const canvas = document.createElement('canvas');
            const gl = canvas.getContext('webgl') || canvas.getContext('experimental-webgl');
            
            if (gl) {
                return [
                    gl.getParameter(gl.VENDOR),
                    gl.getParameter(gl.RENDERER),
                    gl.getParameter(gl.VERSION)
                ].join('|');
            }
        } catch (e) {
            // WebGL不可用
        }
        return 'webgl-unavailable';
    }

    /**
     * 获取字体列表（简化版）
     */
    getFontList() {
        const baseFonts = ['monospace', 'sans-serif', 'serif'];
        const testString = 'mmmmmmmmmmlli';
        const testSize = '72px';
        const h = document.getElementsByTagName('body')[0];
        
        const s = document.createElement('span');
        s.style.fontSize = testSize;
        s.innerHTML = testString;
        h.appendChild(s);
        
        const defaultWidth = {};
        const defaultHeight = {};
        
        for (let i = 0; i < baseFonts.length; i++) {
            s.style.fontFamily = baseFonts[i];
            defaultWidth[baseFonts[i]] = s.offsetWidth;
            defaultHeight[baseFonts[i]] = s.offsetHeight;
        }
        
        h.removeChild(s);
        
        return baseFonts.join(',');
    }

    /**
     * 生成会话ID
     */
    generateSessionId() {
        if (this.sessionId) {
            return this.sessionId;
        }

        // 尝试从localStorage获取
        this.sessionId = localStorage.getItem('browser_session_id');
        
        if (!this.sessionId) {
            // 生成新的会话ID
            this.sessionId = this.generateFingerprint() + '_' + Date.now();
            localStorage.setItem('browser_session_id', this.sessionId);
        }
        
        return this.sessionId;
    }

    /**
     * 简单哈希函数
     */
    simpleHash(str) {
        let hash = 0;
        if (str.length === 0) return hash.toString();
        
        for (let i = 0; i < str.length; i++) {
            const char = str.charCodeAt(i);
            hash = ((hash << 5) - hash) + char;
            hash = hash & hash; // 转换为32位整数
        }
        
        return Math.abs(hash).toString(36);
    }

    /**
     * 获取完整的指纹信息
     */
    getFingerprintInfo() {
        return {
            sessionId: this.generateSessionId(),
            fingerprint: this.generateFingerprint(),
            screenResolution: this.getScreenInfo(),
            timezone: this.getTimeZone(),
            language: this.getLanguage(),
            platform: this.getPlatform(),
            userAgent: navigator.userAgent
        };
    }

    /**
     * 发送指纹信息到服务器
     */
    async sendToServer() {
        const fingerprintInfo = this.getFingerprintInfo();
        
        try {
            const response = await fetch('/api/v1/game-engine/fingerprint', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(fingerprintInfo)
            });
            
            if (response.ok) {
                console.log('指纹信息已发送到服务器');
                return true;
            } else {
                console.error('发送指纹信息失败');
                return false;
            }
        } catch (error) {
            console.error('发送指纹信息时出错:', error);
            return false;
        }
    }
}

// 使用示例
document.addEventListener('DOMContentLoaded', function() {
    const fingerprint = new BrowserFingerprint();
    
    // 生成指纹信息
    const info = fingerprint.getFingerprintInfo();
    console.log('浏览器指纹信息:', info);
    
    // 发送到服务器（可选）
    // fingerprint.sendToServer();
    
    // 将指纹信息存储到全局变量，供其他脚本使用
    window.browserFingerprint = fingerprint;
});

// 导出类（如果使用模块系统）
if (typeof module !== 'undefined' && module.exports) {
    module.exports = BrowserFingerprint;
} 