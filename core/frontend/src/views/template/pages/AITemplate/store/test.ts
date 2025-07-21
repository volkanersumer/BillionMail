export const str = `We are creating a welcome email for BillionMail. The email should be visually appealing and have high conversion.
 We are given company information and style configurations.

 Steps:
 1. Use the provided style config:
    - AccentColor: #f6f6f7 (use as background for some sections, but note: it's very light, almost white)
    - TextColor: #3c3c43 (for text)
    - PageBackground: #ffffff (body background)
    - ContainerBackground: #ffffff (container background)
    - LinkSocialColor: #67676c (for social links, but we might not have social in this email)
    - LinkFooterColor: #20a53a (for unsubscribe and other footer links)
    - Fonts: Inter, but since we are in email, we'll use fallback fonts: 'Inter', Arial, sans-serif

 2. The email must have:
    - A subject: We'll put in the title and also as a hidden preview text.
    - Preheader text (hidden but shown in inbox preview)
    - A header with logo and maybe a welcome message.
    - Body content: Thank the user, highlight key benefits, and include a call to action (like getting started).
    - Footer: Copyright, disclaimer, and unsubscribe link.

 3. We are to use placeholder for unsubscribe: {{ UnsubscribeURL . }}

 4. We have a logo from the company: PrimaryLogo: https://www.billionmail.com/logo.png

 5. We can use emojis for icons.

 6. We are to design for 600px max width.

 7. We have knowledge base content that we can use for the body. We are to create a welcome email, so we don't need to replicate the entire homepage, but we can highlight:

    - Open-source, self-hosted, free from monthly fees.
    - Quick setup (8 minutes to install)
    - Features: Cold Email Partner, Higher Deliverability, Data Visualization Dashboard.

 8. We have images, but we should use sparingly and with alt text. We have:
      Image 1: https://www.billionmail.com/home.png (Alt: BillionMail)
      Image 2: https://img.youtube.com/vi/UHgxZa_9jGs/maxresdefault.jpg (we can use as a thumbnail for a video)

 9. We can include a button for getting started (QuickStart link: /start/quickstart.html -> full URL: https://www.billionmail.com/start/quickstart.html)

 10. We are to use the color scheme: Primary (60%) + Secondary (30%) + Accent (10%).
     Let's define:
        Primary: #20a53a (from LinkFooterColor, which is a green) - we can use for buttons and important highlights.
        Secondary: #3c3c43 (TextColor) for text and some elements.
        Accent: #f6f6f7 (AccentColor) for backgrounds.

     But note: the style config has AccentColor as #f6f6f7 and TextColor as #3c3c43. We also have LinkFooterColor #20a53a which is a bright green.

     We'll use:
        Primary: #20a53a (green) for buttons and key highlights (60%)
        Secondary: #3c3c43 (dark gray) for text and headings (30%)
        Accent: #f6f6f7 (very light gray) for backgrounds (10%)

 11. We are to compress CSS into a single line. But for development, we'll write it in a readable way and then minify.

 12. We are to include a hidden preview text.

 Let's structure the email:

    - Preheader (hidden div with preview text)
    - Container (max-width: 600px; margin: auto; background: #ffffff)
    - Header:
          Logo (centered, with link to home)
          Welcome heading: "Welcome to BillionMail!" (using h1)
          Subheading: "Thank you for joining the open-source email revolution"
    - Main content:
          A short paragraph: "You've taken the first step towards owning your email infrastructure. With BillionMail, you get a powerful, self-hosted mail server and email marketing solution — free from monthly fees and fully under your control."
          Key features (using emojis and bold text):
            - 📧 Cold Email Partner: Unlimited email sending, no built-in speed limit.
            - 🛬 Higher Deliverability: Step-by-step guidance for DNS records and real-time analytics.
            - 📊 Data Visualization Dashboard: Real-time metrics for your campaigns.
          Get started button: "Get Started in Minutes" -> links to QuickStart page.
          Image: We can use the home.png image to show the dashboard? (but note: images might be blocked, so we'll set alt text and style accordingly)
    - Secondary content:
          "Just 3 steps to send a billion emails!" (with emoji: 💌)
          Steps (in a table or divs with numbers):
            1. Install BillionMail: Only 8 minutes.
            2. Connect Your Domain: Add domain, verify DNS, enable SSL.
            3. Build Your Campaign: Write email, choose list, send.
          Another button: "Watch on YouTube" (if we have a video) and "Star us on GitHub"
    - Footer:
          Copyright and disclaimer.
          Unsubscribe link.

 13. We'll use tables for layout to ensure email client compatibility.

 14. We'll include a hidden preview text at the top.

 Implementation:

    We'll write the HTML and CSS. We'll use inline styles. We'll minify the CSS at the end.

    We are to output the entire HTML in a code block.

 Let's code accordingly.
`

