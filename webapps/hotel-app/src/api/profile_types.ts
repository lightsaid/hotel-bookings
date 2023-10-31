
export type LoginType = 0 | 1 // 0: 密码登录，1: 验证码等内陆

// 注册入参
export type RegisterRequest = {
    username: string;
    phone_number: string;
    password: string;
    sms_code: string;
}

// 登录入参：密码登录，password 必填，验证码登录 sms_code 必填
export type LoginRequest = {
    phone_number: string;
    password: string;
    sms_code: string;
    login_type: LoginType;
}


// 登录出参(用户个人信息)
export type ProfileType = {
    id: number;
    phone_number: string;
    username: string;
    avatar: string;
    created_at: string;
    updated_at: string;
    access_token: string;
    refresh_token: string;
}
