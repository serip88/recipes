/*-------------------------------------
  Template name:      Hexadash tailwind css
  Version:            1.0
  Author Name:        dashboardmarket
---------------------------------------*/

"use strict";

/* Preloader */
function preloader(param) {
  let preloader = document.querySelector(param);
  window.addEventListener('load', function () {
    if (preloader) {
      preloader.classList.add("show")
    }
  });
};

/* Search Dropdown trigger */
function toggleDropdown(selector, target) {
  const searchToggle = document.querySelector(selector);
  const searchFormTopMenu = document.querySelector(target);
  if (searchToggle) {
    searchToggle.addEventListener('click', function () {
      this.classList.toggle('close');
      searchFormTopMenu.classList.toggle('visible');
    });
  }
};

/* Dark Switch Button */
function enableDarkMode() {
  const isDarkModeEnabled = localStorage.getItem('darkMode') === 'true';
  const darkModeToggle = document.querySelector('.dark-single-switch');
  const darkModeSwitch = document.querySelector('#darkModeToggle');
  const darkMode = document.querySelector('body');

  function setDarkMode(enabled) {
    if (enabled) {
      darkMode.classList.add('dark');
      localStorage.setItem('darkMode', 'true');
      if (darkModeSwitch) {
        darkModeSwitch.checked = true;
      }
    } else {
      darkMode.classList.remove('dark');
      localStorage.setItem('darkMode', 'false');
      if (darkModeSwitch) {
        darkModeSwitch.checked = false;
      }
    }
  }

  function toggleDarkMode() {
    const darkModeState = darkMode.classList.contains('dark');

    setDarkMode(!darkModeState);
    if (darkModeToggle) {
      darkModeToggle.classList.toggle('active', !darkModeState);
    }
  }

  setDarkMode(isDarkModeEnabled);

  if (darkModeToggle) {
    darkModeToggle.addEventListener('click', toggleDarkMode);
  }

  if (darkModeSwitch) {
    darkModeSwitch.addEventListener('change', toggleDarkMode);
  }
}

/* RTL LTR Switch */
function enableRtlMode() {
  const rtlToggle = document.getElementById('rtlToggle');

  if (rtlToggle) {
    function setDirection() {
      const isRtl = rtlToggle.checked;
      document.documentElement.dir = isRtl ? 'rtl' : 'ltr';
      localStorage.setItem('isRtl', isRtl); // Store the state
    }

    rtlToggle.addEventListener('change', setDirection);

    const isRtlStored = localStorage.getItem('isRtl');
    if (isRtlStored === 'true') {
      rtlToggle.checked = true;
    } else {
      rtlToggle.checked = false;
    }

    setDirection();
  }
}

/* Top Menu Toggle */
function toggleTopMenu() {
  const topMenuToggle = document.getElementById('topMenuToggle');
  const target = document.getElementById('asideBar');
  const target2 = document.getElementById('navBar');
  const body = document.getElementById('content');
  const limToggler = document.getElementById('slim-toggler');
  const topMenuWrapper = document.getElementById('topMenuWrapper');
  const topMenuLogo = document.getElementById('topMenu-logo');
  if (topMenuToggle && target && body && limToggler && topMenuWrapper && topMenuLogo) {
    function setTopMenuStyle() {
      target.classList.toggle('TopCollapsed')
      target2.classList.toggle('TopCollapsed')
      body.classList.toggle('TopExpanded')
      limToggler.classList.toggle('hide', topMenuToggle.checked);
      topMenuWrapper.classList.toggle('flex', topMenuToggle.checked);
      topMenuWrapper.style.transition = topMenuToggle.checked ? 'all 0.3s linear 0s' : 'all 0.3s linear 0s';
      topMenuLogo.classList.toggle('flex', topMenuToggle.checked);
      topMenuLogo.style.transition = topMenuToggle.checked ? 'all 0.3s linear 0s' : 'all 0.3s linear 0s';
    }

    topMenuToggle.addEventListener('change', setTopMenuStyle);
  }
};

/* sidebar dropdown */
function handleSidebarDropdown() {
  let submenus = document.querySelectorAll('aside .sub-item-wrapper');
  let direction = document.querySelector('html').getAttribute('dir');
  submenus.forEach(item => {
    item.addEventListener('mouseover', function () {
      let menuItem = this;
      let menuItemRect = menuItem.getBoundingClientRect()
      let submenuWrapper = menuItem.querySelector('.sub-item');
      submenuWrapper.style.top = `${menuItemRect.top}px`;
    })
  });
}

/* Password visibility */
function togglePasswordVisibility() {
  const passwordToggle = document.querySelector('.js-password-toggle');
  const passwords = document.querySelectorAll('.js-password');
  const passwordLabel = document.querySelector('.js-password-label i');

  if (passwordToggle) {
    passwordToggle.addEventListener('change', function () {
      passwords.forEach(password => {
        if (password.type === 'password') {
          password.type = 'text';
          passwordLabel.classList.add('uil-eye');
          passwordLabel.classList.remove('uil-eye-slash');
        } else {
          password.type = 'password';
          passwordLabel.classList.remove('uil-eye');
          passwordLabel.classList.add('uil-eye-slash');
        }

        password.focus();
      });
    });
  }
};

/* Redirect index page */
function redirectIndexPage(selector) {
  const form = document.getElementById(selector);
  if (form) {
    form.addEventListener('submit', function (e) {
      e.preventDefault();
      if (validateForm()) {
        window.location.href = 'index.html';
      }
    });

    function validateForm() {
      const passwordField = document.getElementById('password');
      if (
        (passwordField.value.trim() === '')
      ) {
        alert('Password is required.');
        return false;
      }
      return true;
    }
  }
};

/* Current Year */
function displayCurrentYear(selector) {
  const elements = document.querySelectorAll(selector);
  const currentYear = new Date().getFullYear();
  if (currentYear) {
    elements.forEach(element => {
      element.innerHTML = currentYear;
    });
  }
};

/* Slim Header */
function toggleSlimView() {
  const sidebarToggle = document.querySelector("#slim-toggler");
  const sidebar = document.querySelector("#asideBar");
  const navbar = document.querySelector("#navBar");
  const contents = document.querySelector("#content");
  const rightEllipsis = document.querySelector("#right-ellipsis")

  function toggleSidebar() {
    if (sidebar) {
      sidebar.classList.toggle("collapsed");
      navbar.classList.toggle("collapsed");
    }
    if (contents) {
      contents.classList.toggle("expanded");
    }
  }

  if (sidebarToggle) {
    sidebarToggle.addEventListener("click", function (e) {
      toggleSidebar();
    });
  }

  function toggleRightEllipsis() {
    if (rightEllipsis) {
      contents.classList.toggle("active");
    }
  }

  if (rightEllipsis) {
    rightEllipsis.addEventListener("click", function (e) {
      toggleRightEllipsis();
    });
  }

  document.addEventListener("click", function (e) {
    const target = e.target; // Clicked element
    if (window.innerWidth <= 1199 && !sidebar.contains(target) && !sidebarToggle.contains(target)) {
      // Check if the sidebar or toggle button was not clicked
      if (sidebar.classList.contains("collapsed")) {
        // If sidebar is collapsed, remove the classes
        sidebar.classList.remove("collapsed");
        navbar.classList.remove("collapsed");
        if (contents) {
          contents.classList.remove("expanded");
        }
      }
    }
  });
};

/* Sidebar Toggle ul */
function toggleSidebarFull() {
  document.querySelectorAll(".navBar .sub-item-wrapper > a").forEach(function (element) {
    element.addEventListener("click", function (e) {
      e.preventDefault();
      var parent = this.parentNode;
      var siblingUl = parent.querySelector("ul");

      document.querySelectorAll(".navBar .sub-item-wrapper").forEach(function (item) {
        if (item !== parent) {
          item.classList.remove("open");
          item.querySelector("ul").style.display = "none";
        }
      });

      parent.classList.toggle("open");
      siblingUl.style.display = parent.classList.contains("open") ? "block" : "none";
    });
  });
}

/* Sidebar scroll sync menu item */
function highlightActiveMenuItem(selector, target) {

  var activeMenuItems = document.querySelectorAll(selector);
  var navBar = document.querySelector(target);

  if (activeMenuItems) {
    activeMenuItems.forEach(function (menuItem) {
      var rect = menuItem.getBoundingClientRect();
      var topValue = rect.top;

      function scrollWin(topValue) {
        if (navBar) {
          var scrollPosition = (topValue - 300);
          navBar.scrollTop = scrollPosition;
        }
      }

      window.onload = function () {
        scrollWin(topValue);
      };
    });
  }
};

/* list Item active class */
function highlightActiveListItem(selector, toggleClass) {
  const links = document.querySelectorAll(selector);

  links.forEach((link) => {
    link.addEventListener("click", function () {
      links.forEach((e) => {
        e.classList.remove(toggleClass)
      })
      this.classList.add(toggleClass)
    })
  })
};

/* Inbox utility */
function handleInboxStars(target) {
  //toggle star
  const toggleStars = document.querySelectorAll(target);
  toggleStars.forEach((toggleStar) => {
    toggleStar.addEventListener("click", function () {
      if (toggleStar.classList.contains("active")) {
        toggleStar.classList.remove("active");
      } else {
        toggleStar.classList.add("active");
      }
    });

  });

  /* Add Label */
  function addNewLabel() {
    const labelName = document.getElementById('newLabelInput').value;
    if (labelName) {
      const listItem = document.createElement('li');
      const anchor = document.createElement('a');
      anchor.href = '#';
      anchor.className = 'font-normal text-[15px] relative flex items-center bg-transparent text-body dark:text-subtitle-dark px-[15px] py-[10px] gap-[15px] rounded-md capitalize hover-bg-primary/10 hover-text-primary';
      anchor.innerHTML = `
                <i class="uil uil-list-ul"></i>
                ${labelName}
            `;
      listItem.appendChild(anchor);
      const labelList = document.getElementById('label-list');
      labelList.appendChild(listItem);
      document.getElementById('newLabelInput').value = ''; // Clear the input field
    }
  }
  const addLabelButton = document.getElementById('addLabelButton');
  if (addLabelButton) {
    addLabelButton.addEventListener('click', addNewLabel);
  }

};

/* Show and Hide Div */
function toggleEmailOverlay(a, b) {
  const replyResult = document.querySelector(a);
  if (replyResult) {
    replyResult.addEventListener('click', function () {
      var div = document.querySelector(b);
      if (div.classList.contains('hidden')) {
        div.classList.remove('hidden');
        div.classList.add('visible');
        replyResult.classList.add('active');
      } else {
        replyResult.classList.remove('active');
        div.classList.remove('visible');
        div.classList.add('hidden');
      }
    });
  }
};

/* Toggle Class */
function adjustWindowMode(selector, targetSelector, toggleClass) {
  const myButtons = document.querySelectorAll(selector);
  const targetDiv = document.querySelector(targetSelector);

  if (myButtons && targetDiv) {
    myButtons.forEach(myButton => {
      myButton.addEventListener("click", () => {
        targetDiv.classList.toggle(toggleClass);
      });
    });
  }
};

/* Checkbox Group Toggle */
function toggleCheckboxGroup(target) {
  const checkboxes = document.querySelectorAll(target);
  checkboxes.forEach((checkbox) => {
    checkbox.addEventListener('click', function () {
      if (this.checked) {
        checkboxes.forEach((otherCheckbox) => {
          if (otherCheckbox !== this) {
            otherCheckbox.checked = false;
          }
        });
      }
    });
  });
};

/* Switch toggle class */
function toggleSwitch(selector, target, class1) {
  const switchSelector = document.getElementById(selector);
  const switchOn = document.getElementById(target);

  if (switchSelector && switchOn) {
    function applyToggleClass() {
      if (switchSelector.checked) {
        switchOn.classList.add(class1);
      } else {
        switchOn.classList.remove(class1);
      }
    }
    switchSelector.addEventListener('change', applyToggleClass);
  }
}

/* Outside Toggle Class */
function toggleClassOutsideClick(target, selector, toggleClass) {
  var button = document.getElementById(target);
  var ul = document.getElementById(selector);

  if (button) {
    button.addEventListener('click', function () {
      ul.classList.toggle(toggleClass);
      button.classList.toggle('active');
    });

    document.addEventListener('click', function (event) {
      var isClickInsideButton = button.contains(event.target);
      var isClickInsideUl = ul.contains(event.target);

      if (!isClickInsideButton && !isClickInsideUl) {
        ul.classList.remove(toggleClass);
        button.classList.remove('active');
      }
    });
  }
};

/* Inline SVG */
function insertInlineSvg(selector) {
  const svg = document.querySelectorAll(selector);
  if (svg) {
    svg.forEach(function (element) {
      var imgID = element.id;
      var imgClass = element.className;
      var imgURL = element.src;

      fetch(imgURL)
        .then(response => response.text())
        .then(function (data) {
          var parser = new DOMParser();
          var xmlDoc = parser.parseFromString(data, 'image/svg+xml');
          var $svg = xmlDoc.getElementsByTagName('svg')[0];

          if (typeof imgID !== 'undefined' && imgID !== '') {
            $svg.setAttribute('id', imgID);
          }

          if (typeof imgClass !== 'undefined' && imgClass !== '') {
            $svg.setAttribute('class', imgClass + ' replaced-svg');
          }

          $svg.removeAttribute('xmlns:a');
          element.parentNode.replaceChild($svg, element);
        });
    });
  }
};

function content_ready_scripts() {
  preloader('.preloader');
  toggleDropdown('.theme-dropdown-trigger', '.theme-dropdown');
  enableDarkMode();
  enableRtlMode();
  toggleTopMenu();
  handleSidebarDropdown();
  togglePasswordVisibility();
  redirectIndexPage('admin-form');
  displayCurrentYear('.current-year');
  toggleSlimView();
  toggleSidebarFull();
  highlightActiveMenuItem('.navBar li.open', '.navBar');
  highlightActiveListItem('.listItemActive li a', 'active');
  handleInboxStars(".inboxToggleStar");
  toggleEmailOverlay('.user-reply-1', '.user-reply-1-result');
  toggleEmailOverlay('.user-reply-2', '.user-reply-2-result');
  toggleEmailOverlay('.email-overlay-btn', '.email-overlay-message');
  toggleEmailOverlay('.email-overlay-message-close', '.email-overlay-message');
  adjustWindowMode('.full-window-trigger', '.email-overlay-message', 'window-full');
  toggleCheckboxGroup('.radio-group input');
  toggleSwitch('switch-selector', 'switch-on', 'active');
  toggleClassOutsideClick('author-dropdown', 'right-ellipsis-trigger', 'active');
  toggleClassOutsideClick('inbox-sidebar-selector', 'inbox-sidebar-target', 'nav-open');
  toggleClassOutsideClick('chat-sidebar-selector', 'chat-sidebar-target', 'nav-open');
  toggleClassOutsideClick('faq-sidebar-selector', 'faq-sidebar-target', 'nav-open');
  insertInlineSvg('img.svg');
  handleInboxStars(".toggle-active");
  highlightActiveMenuItem('.chat-last-message', '.chat-wrapper');
}

setTimeout(function () {

}, 2000);

window.addEventListener('resize', function () {

});

document.addEventListener('DOMContentLoaded', function () {
  content_ready_scripts();
});