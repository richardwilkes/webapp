using System;
using System.Collections.Generic;
using System.Drawing;
using System.Runtime.InteropServices;
using System.Windows.Forms;
using Microsoft.Toolkit.Forms.UI.Controls;

namespace WebApp
{
    public delegate void VoidCallback();
    public delegate byte BoolCallback();
    public delegate int Int32Callback();
    public delegate void VoidPointerCallback(IntPtr arg0);

    public class PlatformWindows
    {
        public const int QuitResponseCancel = 0;
        public const int QuitResponseNow = 1;
        public const int QuitResponseLater = 2;

        private readonly VoidCallback willFinishStartup;
        private readonly VoidCallback didFinishStartup;
        private readonly VoidCallback willActivate;
        private readonly VoidCallback didActivate;
        private readonly VoidCallback willDeactivate;
        private readonly VoidCallback didDeactivate;
        private readonly BoolCallback quitAfterLastWindowClosed;
        private readonly Int32Callback checkQuit;
        private readonly VoidPointerCallback handleMenuItem;

        private RootWindow rootWindow;
        private List<WebWindow> webWindows = new List<WebWindow>();
        private MenuBar menuBar;

        public PlatformWindows(IntPtr willFinishStartup, IntPtr didFinishStartup, IntPtr willActivate, IntPtr didActivate, IntPtr willDeactivate, IntPtr didDeactivate, IntPtr quitAfterLastWindowClosed, IntPtr checkQuit, IntPtr handleMenuItem)
        {
            this.willFinishStartup = Marshal.GetDelegateForFunctionPointer<VoidCallback>(willFinishStartup);
            this.didFinishStartup = Marshal.GetDelegateForFunctionPointer<VoidCallback>(didFinishStartup);
            this.willActivate = Marshal.GetDelegateForFunctionPointer<VoidCallback>(willActivate);
            this.didActivate = Marshal.GetDelegateForFunctionPointer<VoidCallback>(didActivate);
            this.willDeactivate = Marshal.GetDelegateForFunctionPointer<VoidCallback>(willDeactivate);
            this.didDeactivate = Marshal.GetDelegateForFunctionPointer<VoidCallback>(didDeactivate);
            this.quitAfterLastWindowClosed = Marshal.GetDelegateForFunctionPointer<BoolCallback>(quitAfterLastWindowClosed);
            this.checkQuit = Marshal.GetDelegateForFunctionPointer<Int32Callback>(checkQuit);
            this.handleMenuItem = Marshal.GetDelegateForFunctionPointer<VoidPointerCallback>(handleMenuItem);
        }

        public void Start()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            rootWindow = new RootWindow(checkQuit);
            willFinishStartup();
            didFinishStartup();
            // DR Is this the right place to call these?
            willActivate();
            didActivate();
            Application.Run(rootWindow);
        }

        public void SetMenuBar(MenuBar menuBar)
        {
            this.menuBar = menuBar;
            rootWindow.SetMenuBar(menuBar);
            foreach (var webWindow in webWindows)
            {
                webWindow.SetMenuBar(menuBar);
            }
        }

        public WebWindow NewWindow(int width, int height, String url)
        {
            WebWindow WebWindow = new WebWindow(width, height, url);
            webWindows.Add(WebWindow);
            WebWindow.SetMenuBar(menuBar);
            return WebWindow;
        }

        public Menu NewMenuItem(String title)
        {
            return new Menu(this, title);
        }

        public void HandleMenuItem(IntPtr goHandle)
        {
            handleMenuItem(goHandle);
        }
    }

    public class RootWindow : Form
    {
        private readonly Int32Callback checkQuit;

        public RootWindow(Int32Callback checkQuit)
        {
            this.checkQuit = checkQuit;
            FormClosing += new FormClosingEventHandler(OnFormClosing);
        }

        public void SetMenuBar(MenuBar menu)
        {
            Controls.Add(menu);
            MainMenuStrip = menu;
        }

        private void OnFormClosing(object sender, FormClosingEventArgs e)
        {
            switch (checkQuit())
            {
                case PlatformWindows.QuitResponseCancel:
                case PlatformWindows.QuitResponseLater:
                    e.Cancel = true;
                    break;
                case PlatformWindows.QuitResponseNow:
                    // Do nothing
                    break;
            }
        }
    }

    public class WebWindow : Form
    {
        private readonly WebView webView;

        public WebWindow(int width, int height, String url)
        {
            webView = new WebView();
            ((System.ComponentModel.ISupportInitialize)(webView)).BeginInit();
            webView.Dock = DockStyle.Fill;
            Controls.Add(webView);
            ((System.ComponentModel.ISupportInitialize)(webView)).EndInit();
            ClientSize = new Size(width, height);
            webView.Navigate(url);
            Show();
        }

        public void SetMenuBar(MenuBar menu)
        {
            Controls.Add(menu);
            MainMenuStrip = menu;
        }
    }

    public interface MenuFuncs {
        void InsertItem(ToolStripItem child, int index);
        int GetCount();
    }

    public class MenuBar : MenuStrip, MenuFuncs
    {
        public void InsertItem(ToolStripItem child, int index)
        {
            Items.Insert(index, child);
        }

        public int GetCount()
        {
            return Items.Count;
        }
    }

    public class Menu : ToolStripMenuItem, MenuFuncs
    {
        private readonly PlatformWindows platformWindows;
        public IntPtr GoHandle { get; set; }

        public Menu(PlatformWindows platformWindows, String title)
        {
            this.platformWindows = platformWindows;
            Text = title;
            if (platformWindows != null)
            {
                Click += new System.EventHandler(OnClick);
            }
        }

        public void InsertItem(ToolStripItem child, int index)
        {
            DropDownItems.Insert(index, child);
        }

        public int GetCount()
        {
            return DropDownItems.Count;
        }

        private void OnClick(object sender, EventArgs e)
        {
            platformWindows.HandleMenuItem(GoHandle);
        }
    }

    public class MenuItemSeparator : ToolStripSeparator
    {
    }
}
