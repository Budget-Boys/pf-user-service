<?php

namespace App\Http\Controllers;

use App\Http\Requests\UserRequest;
use Illuminate\Http\Request;

class UserController extends Controller
{
    public function store(Request $request)
    {
        return response()->json(['ok' => true]);
    }
}
