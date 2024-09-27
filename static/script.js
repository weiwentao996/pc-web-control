const joystick = document.getElementById('joystick');
const stick = document.getElementById('stick');

let isMoving = false;

document.addEventListener('touchmove', function(e) {
    e.preventDefault(); // 禁止页面拖动
}, { passive: false });

stick.addEventListener('touchstart', (e) => {
    isMoving = true;
    moveStick(e.touches[0]);
});

document.addEventListener('touchend', () => {
    isMoving = false;
    stick.style.transform = 'translate(-50%, -50%)'; // 归位
});

joystick.addEventListener('touchmove', (e) => {
    if (isMoving) {
        moveStick(e.touches[0]);
    }
});

function moveStick(touch) {
    const rect = joystick.getBoundingClientRect();
    const joystickCenterX = rect.left + rect.width / 2;
    const joystickCenterY = rect.top + rect.height / 2;

    const deltaX = touch.clientX - joystickCenterX;
    const deltaY = touch.clientY - joystickCenterY;

    const distance = Math.min(Math.sqrt(deltaX * deltaX + deltaY * deltaY), 100); // 最大移动距离
    const angle = Math.atan2(deltaY, deltaX);

    const stickX = distance * Math.cos(angle);
    const stickY = distance * Math.sin(angle);

    // 保持控制杆在圆形内
    stick.style.transform = `translate(${stickX}px, ${stickY}px) translate(-50%, -50%)`;

    fetch('/move?x=' + Math.floor(stickX*0.1) + '&y=' + Math.floor(stickY*0.1));
}

