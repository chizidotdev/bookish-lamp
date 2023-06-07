import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

// This function can be marked `async` if using `await` inside
export function middleware(request: NextRequest) {
    const authCookie = request.cookies.get('Authorization');
    if (!authCookie) {
        // localStorage.removeItem('token');
        return NextResponse.next();
    }

    // localStorage.setItem('token', authCookie.value);

    return NextResponse.next();
}
